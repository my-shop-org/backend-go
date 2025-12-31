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

type AttributeRepository struct {
	db *gorm.DB
}

func NewAttributeRepository(db *gorm.DB) *AttributeRepository {
	return &AttributeRepository{db: db}
}

func (r *AttributeRepository) CheckIfAttributeExists(id string) bool {
	var a entity.Attribute
	err := r.db.First(&a, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		return false
	}
	return true
}

func (r *AttributeRepository) GetAllAttributes() ([]response.AttributeResponse, error) {
	var attrs []entity.Attribute
	if err := r.db.Preload("Values").Find(&attrs).Error; err != nil {
		return nil, err
	}

	res := make([]response.AttributeResponse, 0, len(attrs))
	for _, a := range attrs {
		ar := response.AttributeResponse{
			ID:   a.ID,
			Name: a.Name,
		}
		for _, v := range a.Values {
			ar.Values = append(ar.Values, response.AttributeValueResponse{ID: v.ID, Value: v.Value})
		}
		res = append(res, ar)
	}
	if res == nil {
		res = make([]response.AttributeResponse, 0)
	}
	return res, nil
}

func (r *AttributeRepository) GetAttributeByID(id string) (*response.AttributeResponse, error) {
	var a entity.Attribute
	if err := r.db.Preload("Values").First(&a, "id = ?", id).Error; err != nil {
		return nil, err
	}

	ar := &response.AttributeResponse{
		ID:   a.ID,
		Name: a.Name,
	}
	for _, v := range a.Values {
		ar.Values = append(ar.Values, response.AttributeValueResponse{ID: v.ID, Value: v.Value})
	}
	return ar, nil
}

func (r *AttributeRepository) AddAttribute(attr *request.AttributeRequest) error {
	a := entity.Attribute{Name: pkg.CapitalizeFirstLetter(attr.Name)}
	if err := r.db.Create(&a).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return pkg.DuplicateEntry
		}
		return err
	}
	return nil
}

func (r *AttributeRepository) UpdateAttribute(id string, attr *request.AttributePatchRequest) (*response.AttributeResponse, error) {
	if !r.CheckIfAttributeExists(id) {
		return nil, pkg.AttributeNotFound
	}

	updates := map[string]interface{}{}
	if attr.Name != nil {
		updates["name"] = *attr.Name
	}

	if len(updates) == 0 {
		return nil, pkg.NoFieldsToUpdate
	}

	if err := r.db.Model(&entity.Attribute{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, pkg.DuplicateEntry
		}
		return nil, err
	}

	return r.GetAttributeByID(id)
}

func (r *AttributeRepository) DeleteAttribute(id string) error {
	if !r.CheckIfAttributeExists(id) {
		return pkg.AttributeNotFound
	}

	var cnt int64
	if err := r.db.Model(&entity.AttributeValue{}).Where("attribute_id = ?", id).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt > 0 {
		return pkg.AttributeHasValues
	}

	return r.db.Delete(&entity.Attribute{}, "id = ?", id).Error
}
