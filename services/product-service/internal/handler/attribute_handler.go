package handler

import (
	"errors"
	"net/http"
	"product-service/internal/pkg"
	"product-service/internal/request"
	"product-service/internal/usecase"

	"github.com/labstack/echo/v4"
)

type AttributeHandler struct {
	attributeUsecase *usecase.AttributeUsecase
}

func NewAttributeHandler(attributeUsecase *usecase.AttributeUsecase) *AttributeHandler {
	return &AttributeHandler{attributeUsecase: attributeUsecase}
}

func (h *AttributeHandler) GetAllAttributes(c echo.Context) error {
	attrs, err := h.attributeUsecase.GetAllAttributes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve attributes"})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": attrs})
}

func (h *AttributeHandler) AddAttribute(c echo.Context, attr *request.AttributeRequest) error {
	if err := h.attributeUsecase.AddAttribute(attr); err != nil {
		switch {
		case errors.Is(err, pkg.DuplicateEntry):
			return c.JSON(http.StatusConflict, echo.Map{"message": "Attribute name already exists"})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to create attribute"})
		}
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "Attribute created successfully", "data": attr})
}

func (h *AttributeHandler) GetAttributeByID(c echo.Context) error {
	id := c.Param("id")
	attr, err := h.attributeUsecase.GetAttributeByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, attr)
}

func (h *AttributeHandler) PatchAttribute(c echo.Context, attr *request.AttributePatchRequest) error {
	id := c.Param("id")
	updated, err := h.attributeUsecase.UpdateAttribute(id, attr)
	if err != nil {
		switch {
		case errors.Is(err, pkg.AttributeNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Attribute not found"})
		case errors.Is(err, pkg.DuplicateEntry):
			return c.JSON(http.StatusConflict, echo.Map{"message": "Attribute name already exists"})
		case errors.Is(err, pkg.NoFieldsToUpdate):
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "No fields provided to update"})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to update attribute"})
		}
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Attribute updated successfully", "data": updated})
}

func (h *AttributeHandler) DeleteAttribute(c echo.Context) error {
	id := c.Param("id")
	if err := h.attributeUsecase.DeleteAttribute(id); err != nil {
		switch {
		case errors.Is(err, pkg.AttributeNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Attribute not found"})
		case errors.Is(err, pkg.AttributeHasValues):
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "Attribute has values and cannot be deleted"})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
		}
	}
	return c.JSON(http.StatusNoContent, nil)
}
