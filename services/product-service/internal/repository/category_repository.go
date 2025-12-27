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

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategories() ([]response.CategoryResponse, error) {
	var categories []response.CategoryResponse
	if err := r.db.Model(&entity.Category{}).
		Select("id, name, description, parent_id").
		Scan(&categories).Error; err != nil {
		return nil, err
	}
	if categories == nil {
		categories = make([]response.CategoryResponse, 0)
	}
	return categories, nil
}

func (r *CategoryRepository) CheckIfCategoryExists(id string) bool {
	var category entity.Category
	err := r.db.First(&category, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		// For other errors, log or handle as needed. Here, treat as not found.
		return false
	}
	return true
}

func (r *CategoryRepository) GetCategoryByID(id string) (*response.CategoryResponse, error) {
	var category entity.Category
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &response.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ParentID:    category.ParentID,
	}, nil
}

func (r *CategoryRepository) GetChildCategoriesByID(id string) ([]response.CategoryResponse, error) {
	var categories []response.CategoryResponse
	if err := r.db.Model(&entity.Category{}).
		Select("id, name, description, parent_id").
		Where("parent_id = ?", id).
		Scan(&categories).Error; err != nil {
		return nil, err
	}
	if categories == nil {
		categories = make([]response.CategoryResponse, 0)
	}
	return categories, nil
}

func (r *CategoryRepository) AddCategory(category *request.CategoryRequest) error {
	if category.ParentID != nil {
		if !r.CheckIfCategoryExists(pkg.UintToString(*category.ParentID)) {
			return pkg.ParentCategoryNotFound
		}
	}
	cat := entity.Category{
		Name:        category.Name,
		Description: category.Description,
		ParentID:    category.ParentID,
	}
	err := r.db.Create(&cat).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return pkg.DuplicateEntry
		}
		return err
	}
	return nil
}

func (r *CategoryRepository) UpdateCategory(id string, category *request.CategoryPatchRequest) (
	*response.CategoryResponse, error) {
	if !r.CheckIfCategoryExists(id) {
		return nil, pkg.CategoryNotFound
	}

	if category.ParentID != nil {
		if *category.ParentID == pkg.StringToUint(id) {
			return nil, pkg.CategoryCannotBeItsOwnParent
		}
		if !r.CheckIfCategoryExists(pkg.UintToString(*category.ParentID)) {
			return nil, pkg.ParentCategoryNotFound
		}
	}

	if err := r.db.Model(&entity.Category{}).Where("id = ?", id).Updates(category).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, pkg.DuplicateEntry
		}
		return nil, err
	}

	return nil, nil
}

func (r *CategoryRepository) DeleteCategory(id string) error {
	if !r.CheckIfCategoryExists(id) {
		return pkg.CategoryNotFound
	}
	// Check if category is parent of any child category
	var childCount int64
	if err := r.db.Model(&entity.Category{}).Where("parent_id = ?", id).Count(&childCount).Error; err != nil {
		return err
	}

	if childCount > 0 {
		return pkg.CategoryHasChildren // You should define this error in your pkg
	}

	return r.db.Delete(&entity.Category{}, "id = ?", id).Error
}

func (r *CategoryRepository) GetCategoryTree() ([]*response.CategoryTreeResponse, error) {
	return r.buildCategoryTree(nil)
}

func (r *CategoryRepository) buildCategoryTree(parentID *uint) ([]*response.CategoryTreeResponse, error) {
	var flatCats []response.CategoryResponse
	if parentID == nil {
		err := r.db.Model(&entity.Category{}).
			Select("id, name, description, parent_id").
			Where("parent_id IS NULL").
			Scan(&flatCats).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := r.db.Model(&entity.Category{}).
			Select("id, name, description, parent_id").
			Where("parent_id = ?", parentID).
			Scan(&flatCats).Error
		if err != nil {
			return nil, err
		}
	}

	var treeCats []*response.CategoryTreeResponse
	for i := range flatCats {
		children, _ := r.buildCategoryTree(&flatCats[i].ID)
		treeCat := &response.CategoryTreeResponse{
			CategoryResponse: flatCats[i],
			Children:         children,
		}
		treeCats = append(treeCats, treeCat)
	}
	return treeCats, nil
}
