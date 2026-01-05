package handler

import (
	"errors"
	"net/http"
	"product-service/internal/request"
	"product-service/internal/response"
	"product-service/internal/usecase"

	"myshop-shared/pkg"

	"github.com/labstack/echo/v4"
)

type ProductImageHandler struct {
	productImageUsecase *usecase.ProductImageUsecase
}

func NewProductImageHandler(productImageUsecase *usecase.ProductImageUsecase) *ProductImageHandler {
	return &ProductImageHandler{productImageUsecase: productImageUsecase}
}

func (h *ProductImageHandler) GetAllProductImages(c echo.Context) error {
	images, err := h.productImageUsecase.GetAllProductImages(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve product images"})
	}

	responses := make([]response.ProductImageResponse, len(images))
	for i, img := range images {
		responses[i] = response.ProductImageResponse{
			ID:        img.ID,
			ProductID: img.ProductID,
			VariantID: img.VariantID,
			URL:       img.URL,
			IsDefault: img.IsDefault,
			CreatedAt: img.CreatedAt,
			UpdatedAt: img.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"data": responses})
}

func (h *ProductImageHandler) GetProductImageByID(c echo.Context) error {
	id := c.Param("id")
	image, err := h.productImageUsecase.GetProductImageByID(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, pkg.ProductImageNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Product image not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve product image"})
	}

	resp := response.ProductImageResponse{
		ID:        image.ID,
		ProductID: image.ProductID,
		VariantID: image.VariantID,
		URL:       image.URL,
		IsDefault: image.IsDefault,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}

	return c.JSON(http.StatusOK, echo.Map{"data": resp})
}

func (h *ProductImageHandler) GetImagesByProductID(c echo.Context) error {
	productID := c.Param("productId")

	id := pkg.StringToUint(productID)
	if id == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid product ID"})
	}

	images, err := h.productImageUsecase.GetImagesByProductID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve product images"})
	}

	responses := make([]response.ProductImageResponse, len(images))
	for i, img := range images {
		responses[i] = response.ProductImageResponse{
			ID:        img.ID,
			ProductID: img.ProductID,
			VariantID: img.VariantID,
			URL:       img.URL,
			IsDefault: img.IsDefault,
			CreatedAt: img.CreatedAt,
			UpdatedAt: img.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"data": responses})
}

func (h *ProductImageHandler) GetImagesByVariantID(c echo.Context) error {
	variantID := c.Param("variantId")

	id := pkg.StringToUint(variantID)
	if id == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid variant ID"})
	}

	images, err := h.productImageUsecase.GetImagesByVariantID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve variant images"})
	}

	responses := make([]response.ProductImageResponse, len(images))
	for i, img := range images {
		responses[i] = response.ProductImageResponse{
			ID:        img.ID,
			ProductID: img.ProductID,
			VariantID: img.VariantID,
			URL:       img.URL,
			IsDefault: img.IsDefault,
			CreatedAt: img.CreatedAt,
			UpdatedAt: img.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"data": responses})
}

func (h *ProductImageHandler) AddProductImage(c echo.Context, imageReq *request.ProductImageRequest) error {
	image, err := h.productImageUsecase.AddProductImage(c.Request().Context(), imageReq)
	if err != nil {
		if errors.Is(err, pkg.ProductNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
		}
		if errors.Is(err, pkg.VariantNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Variant not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to add product image"})
	}

	resp := response.ProductImageResponse{
		ID:        image.ID,
		ProductID: image.ProductID,
		VariantID: image.VariantID,
		URL:       image.URL,
		IsDefault: image.IsDefault,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}

	return c.JSON(http.StatusCreated, echo.Map{"data": resp})
}

func (h *ProductImageHandler) UpdateProductImage(c echo.Context, imageReq *request.ProductImagePatchRequest) error {
	id := c.Param("id")

	image, err := h.productImageUsecase.UpdateProductImage(c.Request().Context(), id, imageReq)
	if err != nil {
		if errors.Is(err, pkg.NoFieldsToUpdate) {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "No fields to update"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to update product image"})
	}

	resp := response.ProductImageResponse{
		ID:        image.ID,
		ProductID: image.ProductID,
		VariantID: image.VariantID,
		URL:       image.URL,
		IsDefault: image.IsDefault,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}

	return c.JSON(http.StatusOK, echo.Map{"data": resp})
}

func (h *ProductImageHandler) DeleteProductImage(c echo.Context) error {
	id := c.Param("id")

	err := h.productImageUsecase.DeleteProductImage(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to delete product image"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Product image deleted successfully"})
}
