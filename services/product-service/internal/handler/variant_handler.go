package handler

import (
	"errors"
	"net/http"
	"product-service/internal/params"
	"product-service/internal/request"
	"product-service/internal/usecase"

	"myshop-shared/pkg"

	"github.com/labstack/echo/v4"
)

type VariantHandler struct {
	variantUsecase *usecase.VariantUsecase
}

func NewVariantHandler(variantUsecase *usecase.VariantUsecase) *VariantHandler {
	return &VariantHandler{variantUsecase: variantUsecase}
}

func (h *VariantHandler) GetAllVariants(c echo.Context) error {
	variantQueryParams := params.NewVariantQueryParam()
	err := c.Bind(variantQueryParams)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid param."})
	}

	variants, err := h.variantUsecase.GetAllVariants(c.Request().Context(), variantQueryParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve variants"})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": variants})
}

func (h *VariantHandler) GetVariantByID(c echo.Context) error {
	id := c.Param("id")
	variant, err := h.variantUsecase.GetVariantByID(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, pkg.VariantNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Variant not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve variant"})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": variant})
}

func (h *VariantHandler) AddVariant(c echo.Context, variant *request.VariantRequest) error {
	if err := h.variantUsecase.AddVariant(c.Request().Context(), variant); err != nil {
		switch {
		case errors.Is(err, pkg.ProductNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
		case errors.Is(err, pkg.AttributeValueNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "One or more attribute values not found"})
		case errors.Is(err, pkg.InvalidAttributeValueForProduct):
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "One or more attribute values do not belong to this product's attributes"})
		case errors.Is(err, pkg.DuplicateEntry):
			return c.JSON(http.StatusConflict, echo.Map{"message": "Variant with this SKU already exists"})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to create variant"})
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Variant created successfully", "data": variant})
}

func (h *VariantHandler) PatchVariant(c echo.Context, variant *request.VariantPatchRequest) error {
	id := c.Param("id")

	updatedVariant, err := h.variantUsecase.UpdateVariant(c.Request().Context(), id, variant)
	if err != nil {
		switch {
		case errors.Is(err, pkg.VariantNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Variant not found"})
		case errors.Is(err, pkg.ProductNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
		case errors.Is(err, pkg.AttributeValueNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "One or more attribute values not found"})
		case errors.Is(err, pkg.InvalidAttributeValueForProduct):
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "One or more attribute values do not belong to this product's attributes"})
		case errors.Is(err, pkg.DuplicateEntry):
			return c.JSON(http.StatusConflict, echo.Map{"message": "Variant with this SKU already exists"})
		case errors.Is(err, pkg.NoFieldsToUpdate):
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "No fields provided to update"})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to update variant"})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Variant updated successfully", "data": updatedVariant})
}

func (h *VariantHandler) DeleteVariant(c echo.Context) error {
	id := c.Param("id")
	if err := h.variantUsecase.DeleteVariant(c.Request().Context(), id); err != nil {
		if errors.Is(err, pkg.VariantNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Variant not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to delete variant"})
	}
	return c.JSON(http.StatusNoContent, nil)
}
