package repository

import (
	"product-service/internal/entity"
	"product-service/internal/pkg"
	"product-service/internal/request"
	"product-service/internal/response"

	"github.com/google/uuid"
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
		ID:          category.ID.String(),
		Name:        category.Name,
		Description: category.Description,
		ParentID:    pkg.UUIDToStringPtr(category.ParentID),
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
		if !r.CheckIfCategoryExists(category.ParentID.String()) {
			return pkg.ParentCategoryNotFound
		}
	}
	cat := entity.Category{
		Name:        category.Name,
		Description: category.Description,
		ParentID:    category.ParentID,
	}
	return r.db.Create(&cat).Error
}

func (r *CategoryRepository) UpdateCategory(id string, category *request.CategoryRequest) (*response.CategoryResponse, error) {
	if category.ParentID != nil && id == category.ParentID.String() {
		return nil, pkg.CategoryCannotBeItsOwnParent
	}
	if !r.CheckIfCategoryExists(id) {
		return nil, pkg.CategoryNotFound
	}
	if category.ParentID != nil && !r.CheckIfCategoryExists(category.ParentID.String()) {
		return nil, pkg.ParentCategoryNotFound
	}
	if err := r.db.Model(&entity.Category{}).Where("id = ?", id).Updates(entity.Category{
		Name:        category.Name,
		Description: category.Description,
		ParentID:    category.ParentID,
	}).Error; err != nil {
		return nil, err
	}
	return &response.CategoryResponse{
		ID:          id,
		Name:        category.Name,
		Description: category.Description,
		ParentID:    pkg.UUIDToStringPtr(category.ParentID),
	}, nil
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

func (r *CategoryRepository) buildCategoryTree(parentID *uuid.UUID) ([]*response.CategoryTreeResponse, error) {
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
		parsedUUID, _ := uuid.Parse(flatCats[i].ID)
		children, _ := r.buildCategoryTree(&parsedUUID)
		treeCat := &response.CategoryTreeResponse{
			CategoryResponse: flatCats[i],
			Children:         children,
		}
		treeCats = append(treeCats, treeCat)
	}
	return treeCats, nil
}
