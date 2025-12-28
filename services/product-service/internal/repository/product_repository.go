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
	// Verify all categories exist
	categories := r.CheckCategoriesExist(productReq.Categories)
	if categories == nil {
		return pkg.CategoryNotFound
	}

	// Use transaction to ensure atomicity
	return r.db.Transaction(func(tx *gorm.DB) error {
		product := entity.Product{
			Name:        productReq.Name,
			Description: productReq.Description,
			Price:       productReq.Price,
		}

		// Create product first
		if err := tx.Create(&product).Error; err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				return pkg.DuplicateEntry
			}
			return err
		}

		// Use Association API to append categories
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

	// if err := r.db.Model(&entity.Product{}).Where("id = ?", id).Updates(productReq).Error; err != nil {
	// 	var pgErr *pgconn.PgError
	// 	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
	// 		return nil, pkg.DuplicateEntry
	// 	}
	// 	return nil, err
	// }

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

func (r *ProductRepository) GetProductsByCategoryID(categoryID string) ([]response.ProductResponse, error) {
	var products []entity.Product
	if err := r.db.Preload("Categories").
		Joins("JOIN product_categories ON products.id = product_categories.product_id").
		Where("product_categories.category_id = ?", categoryID).
		Find(&products).Error; err != nil {
		return nil, err
	}

	productResponses := make([]response.ProductResponse, len(products))
	for i, product := range products {
		categories := make([]response.CategoryResponse, len(product.Categories))
		for j, cat := range product.Categories {
			categories[j] = response.CategoryResponse{
				ID:          cat.ID,
				Name:        cat.Name,
				Description: cat.Description,
				ParentID:    cat.ParentID,
			}
		}

		productResponses[i] = response.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Categories:  categories,
			Price:       product.Price,
		}
	}

	if productResponses == nil {
		productResponses = make([]response.ProductResponse, 0)
	}

	return productResponses, nil
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
