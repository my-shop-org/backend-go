package usecase

import (
	"product-service/internal/repository"
	"product-service/internal/request"
	"product-service/internal/response"
)

type AttributeUsecase struct {
	attributeRepo *repository.AttributeRepository
}

func NewAttributeUsecase(attributeRepo *repository.AttributeRepository) *AttributeUsecase {
	return &AttributeUsecase{attributeRepo: attributeRepo}
}

func (u *AttributeUsecase) GetAllAttributes() ([]response.AttributeResponse, error) {
	return u.attributeRepo.GetAllAttributes()
}

func (u *AttributeUsecase) AddAttribute(attr *request.AttributeRequest) error {
	return u.attributeRepo.AddAttribute(attr)
}

func (u *AttributeUsecase) GetAttributeByID(id string) (*response.AttributeResponse, error) {
	return u.attributeRepo.GetAttributeByID(id)
}

func (u *AttributeUsecase) UpdateAttribute(id string, attr *request.AttributePatchRequest) (*response.AttributeResponse, error) {
	return u.attributeRepo.UpdateAttribute(id, attr)
}

func (u *AttributeUsecase) DeleteAttribute(id string) error {
	return u.attributeRepo.DeleteAttribute(id)
}
