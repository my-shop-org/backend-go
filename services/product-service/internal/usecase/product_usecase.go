package usecase

import (
	"product-service/internal/entity"
	"product-service/internal/repository"
	"product-service/internal/request"
)

type ProductUsecase struct {
	productRepo *repository.ProductRepository
}

func NewProductUsecase(productRepo *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepo: productRepo}
}

func (u *ProductUsecase) GetAllProducts() ([]*entity.Product, error) {
	return u.productRepo.GetAllProducts()
}

func (u *ProductUsecase) GetProductByID(id string) (*entity.Product, error) {
	return u.productRepo.GetProductByID(id)
}

func (u *ProductUsecase) AddProduct(product *request.ProductRequest) error {
	return u.productRepo.AddProduct(product)
}

func (u *ProductUsecase) UpdateProduct(id string, product *request.ProductPatchRequest) (*entity.Product, error) {
	return u.productRepo.UpdateProduct(id, product)
}

func (u *ProductUsecase) DeleteProduct(id string) error {
	return u.productRepo.DeleteProduct(id)
}
