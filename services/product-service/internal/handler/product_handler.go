package handler

import (
	"errors"
	"net/http"
	"product-service/internal/params"
	"product-service/internal/request"
	"product-service/internal/usecase"

	"github.com/kaunghtethein/backend-go/shared/pkg"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}

func NewProductHandler(productUsecase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

func (h *ProductHandler) GetAllProducts(c echo.Context) error {
	productQueryParams := params.NewProductQueryParam()
	err := c.Bind(productQueryParams)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid param."})
	}

	products, err := h.productUsecase.GetAllProducts(c.Request().Context(), productQueryParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve products"})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": products})
}

func (h *ProductHandler) GetProductByID(c echo.Context) error {
	id := c.Param("id")
	product, err := h.productUsecase.GetProductByID(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, pkg.ProductNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve product"})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": product})
}

func (h *ProductHandler) AddProduct(c echo.Context, product *request.ProductRequest) error {
	if err := h.productUsecase.AddProduct(c.Request().Context(), product); err != nil {
		switch {
		case errors.Is(err, pkg.CategoryNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "One or more categories not found"})
		case errors.Is(err, pkg.DuplicateEntry):
			return c.JSON(http.StatusConflict, echo.Map{"message": "Product already exists"})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to create product"})
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Product created successfully", "data": product})
}

func (h *ProductHandler) PatchProduct(c echo.Context, product *request.ProductPatchRequest) error {
	id := c.Param("id")

	updatedProduct, err := h.productUsecase.UpdateProduct(c.Request().Context(), id, product)
	if err != nil {
		switch {
		case errors.Is(err, pkg.ProductNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
		case errors.Is(err, pkg.CategoryNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "One or more categories not found"})
		case errors.Is(err, pkg.DuplicateEntry):
			return c.JSON(http.StatusConflict, echo.Map{"message": "Product name already exists"})
		case errors.Is(err, pkg.NoFieldsToUpdate):
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "No fields provided to update"})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to update product"})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Product updated successfully", "data": updatedProduct})
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	if err := h.productUsecase.DeleteProduct(c.Request().Context(), id); err != nil {
		if errors.Is(err, pkg.ProductNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to delete product"})
	}
	return c.JSON(http.StatusNoContent, nil)
}
