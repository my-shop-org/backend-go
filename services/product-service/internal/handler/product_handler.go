package handler

import (
	"errors"
	"product-service/internal/pkg"
	"product-service/internal/request"
	"product-service/internal/usecase"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}

func NewProductHandler(productUsecase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

func (h *ProductHandler) GetAllProducts(c echo.Context) error {
	products, err := h.productUsecase.GetAllProducts()
	if err != nil {
		return c.JSON(500, echo.Map{"message": "Failed to retrieve products"})
	}
	return c.JSON(200, echo.Map{"data": products})
}

func (h *ProductHandler) GetProductByID(c echo.Context) error {
	id := c.Param("id")
	product, err := h.productUsecase.GetProductByID(id)
	if err != nil {
		if errors.Is(err, pkg.ProductNotFound) {
			return c.JSON(404, echo.Map{"message": "Product not found"})
		}
		return c.JSON(500, echo.Map{"message": "Failed to retrieve product"})
	}
	return c.JSON(200, echo.Map{"data": product})
}

func (h *ProductHandler) AddProduct(c echo.Context, product *request.ProductRequest) error {
	if err := h.productUsecase.AddProduct(product); err != nil {
		switch {
		case errors.Is(err, pkg.CategoryNotFound):
			return c.JSON(404, echo.Map{"message": "One or more categories not found"})
		case errors.Is(err, pkg.DuplicateEntry):
			return c.JSON(409, echo.Map{"message": "Product already exists"})
		default:
			return c.JSON(500, echo.Map{"message": "Failed to create product"})
		}
	}

	return c.JSON(201, echo.Map{"message": "Product created successfully", "data": product})
}

func (h *ProductHandler) PatchProduct(c echo.Context, product *request.ProductPatchRequest) error {
	id := c.Param("id")

	updatedProduct, err := h.productUsecase.UpdateProduct(id, product)
	if err != nil {
		switch {
		case errors.Is(err, pkg.ProductNotFound):
			return c.JSON(404, echo.Map{"message": "Product not found"})
		case errors.Is(err, pkg.CategoryNotFound):
			return c.JSON(404, echo.Map{"message": "One or more categories not found"})
		case errors.Is(err, pkg.DuplicateEntry):
			return c.JSON(409, echo.Map{"message": "Product name already exists"})
		case errors.Is(err, pkg.NoFieldsToUpdate):
			return c.JSON(400, echo.Map{"message": "No fields provided to update"})
		default:
			return c.JSON(500, echo.Map{"message": "Failed to update product"})
		}
	}

	return c.JSON(200, echo.Map{"message": "Product updated successfully", "data": updatedProduct})
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	if err := h.productUsecase.DeleteProduct(id); err != nil {
		if errors.Is(err, pkg.ProductNotFound) {
			return c.JSON(404, echo.Map{"message": "Product not found"})
		}
		return c.JSON(500, echo.Map{"message": "Failed to delete product"})
	}
	return c.JSON(204, nil)
}

func (h *ProductHandler) GetProductsByCategoryID(c echo.Context) error {
	categoryID := c.Param("id")
	products, err := h.productUsecase.GetProductsByCategoryID(categoryID)
	if err != nil {
		return c.JSON(500, echo.Map{"message": "Failed to retrieve products"})
	}
	return c.JSON(200, echo.Map{"data": products})
}
