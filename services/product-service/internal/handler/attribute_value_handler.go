package handler

import (
	"errors"
	"net/http"
	"product-service/internal/request"
	"product-service/internal/usecase"

	"github.com/kaunghtethein/backend-go/shared/pkg"
	"github.com/labstack/echo/v4"
)

type AttributeValueHandler struct {
	attributeValueUsecase *usecase.AttributeValueUsecase
}

func NewAttributeValueHandler(attributeValueUsecase *usecase.AttributeValueUsecase) *AttributeValueHandler {
	return &AttributeValueHandler{attributeValueUsecase: attributeValueUsecase}
}

func (h *AttributeValueHandler) GetAllAttributeValues(c echo.Context) error {
	avs, err := h.attributeValueUsecase.GetAllAttributeValues()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to retrieve attribute values"})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": avs})
}

func (h *AttributeValueHandler) AddAttributeValue(c echo.Context, avReq *request.AttributeValueRequest) error {
	if err := h.attributeValueUsecase.AddAttributeValue(avReq); err != nil {
		switch {
		case errors.Is(err, pkg.AttributeNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Attribute not found"})
		case errors.Is(err, pkg.DuplicateEntry):
			return c.JSON(http.StatusConflict, echo.Map{"message": "Attribute value already exists"})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to create attribute value"})
		}
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "Attribute value created successfully", "data": avReq})
}

func (h *AttributeValueHandler) GetAttributeValueByID(c echo.Context) error {
	id := c.Param("id")
	av, err := h.attributeValueUsecase.GetAttributeValueByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, av)
}

func (h *AttributeValueHandler) PatchAttributeValue(c echo.Context, avReq *request.AttributeValuePatchRequest) error {
	id := c.Param("id")
	updated, err := h.attributeValueUsecase.UpdateAttributeValue(id, avReq)
	if err != nil {
		switch {
		case errors.Is(err, pkg.AttributeValueNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Attribute value not found"})
		case errors.Is(err, pkg.AttributeNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Attribute not found"})
		case errors.Is(err, pkg.DuplicateEntry):
			return c.JSON(http.StatusConflict, echo.Map{"message": "Attribute value already exists"})
		case errors.Is(err, pkg.NoFieldsToUpdate):
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "No fields provided to update"})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to update attribute value"})
		}
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Attribute value updated successfully", "data": updated})
}

func (h *AttributeValueHandler) DeleteAttributeValue(c echo.Context) error {
	id := c.Param("id")
	if err := h.attributeValueUsecase.DeleteAttributeValue(id); err != nil {
		switch {
		case errors.Is(err, pkg.AttributeValueNotFound):
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Attribute value not found"})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
		}
	}
	return c.JSON(http.StatusNoContent, nil)
}
