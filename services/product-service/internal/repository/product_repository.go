package repository

import (
	"errors"
	"product-service/internal/entity"
	"product-service/internal/pkg"
	"product-service/internal/request"
	"product-service/internal/response"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProducts() ([]*entity.Product, error) {
	var products []*entity.Product
	if err := r.db.Preload("Categories").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id string) (*entity.Product, error) {
	var product = new(entity.Product)
	if err := r.db.Preload("Categories").First(&product, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ProductNotFound
		}
		return nil, err
	}

	categories := make([]response.CategoryResponse, len(product.Categories))
	for i, cat := range product.Categories {
		categories[i] = response.CategoryResponse{
			ID:          cat.ID,
			Name:        cat.Name,
			Description: cat.Description,
			ParentID:    cat.ParentID,
		}
	}

	return product, nil
}

func (r *ProductRepository) AddProduct(productReq *request.ProductRequest) error {
	categories := r.CheckCategoriesExist(productReq.Categories)
	if categories == nil {
		return pkg.CategoryNotFound
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		product := entity.Product{
			Name:        productReq.Name,
			Description: productReq.Description,
			Price:       productReq.Price,
		}

		if err := tx.Create(&product).Error; err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				return pkg.DuplicateEntry
			}
			return err
		}

		if err := tx.Model(&product).Association("Categories").Append(categories); err != nil {
			return err
		}

		return nil
	})
}

func (r *ProductRepository) UpdateProduct(id string, productReq *request.ProductPatchRequest) (
	*entity.Product, error) {
	if productReq.Categories != nil {
		categories := r.CheckCategoriesExist(*productReq.Categories)
		if categories == nil {
			return nil, pkg.CategoryNotFound
		}
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		updates := make(map[string]interface{})
		if productReq.Name != nil {
			updates["name"] = *productReq.Name
		}
		if productReq.Description != nil {
			updates["description"] = *productReq.Description
		}
		if productReq.Price != nil {
			updates["price"] = *productReq.Price
		}

		if productReq.Categories != nil {
			categories := r.CheckCategoriesExist(*productReq.Categories)
			var product entity.Product
			if err := tx.First(&product, "id = ?", id).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return pkg.ProductNotFound
				}
				return err
			}

			if err := tx.Model(&product).Association("Categories").Replace(categories); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return r.GetProductByID(id)
}

func (r *ProductRepository) DeleteProduct(id string) error {
	var product entity.Product
	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return pkg.ProductNotFound
		}
		return err
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&product).Association("Categories").Clear(); err != nil {
			return err
		}
		if err := tx.Delete(&product).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *ProductRepository) CheckCategoriesExist(catIDs []uint) []entity.Category {
	var categories []entity.Category
	if err := r.db.Where("id IN ?", catIDs).Find(&categories).Error; err != nil {
		return nil
	}

	if len(categories) != len(catIDs) {
		return nil
	}

	return categories
}
