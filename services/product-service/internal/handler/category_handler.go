package handler

import (
	"errors"
	"product-service/internal/pkg"
	"product-service/internal/request"
	"product-service/internal/usecase"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryUsecase *usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase *usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase: categoryUsecase}
}

func (h *CategoryHandler) GetAllCategories(c echo.Context) error {
	categories, err := h.categoryUsecase.GetAllCategories()
	if err != nil {
		return c.JSON(500, echo.Map{"message": "Failed to retrieve categories"})
	}
	return c.JSON(200, echo.Map{"data": categories})
}

func (h *CategoryHandler) AddCategory(c echo.Context, category *request.CategoryRequest) error {
	if err := h.categoryUsecase.AddCategory(category); err != nil {
		switch {
		case errors.Is(err, pkg.DuplicateEntry):
			return c.JSON(409, echo.Map{"message": "Category name already exists"})
		case errors.Is(err, pkg.ParentCategoryNotFound):
			return c.JSON(404, echo.Map{"message": "Parent category not found"})
		default:
			return c.JSON(500, echo.Map{"message": "Failed to create category"})
		}
	}

	return c.JSON(201, echo.Map{"message": "Category created successfully", "data": category})
}

func (h *CategoryHandler) GetCategoryByID(c echo.Context) error {
	id := c.Param("id")
	category, err := h.categoryUsecase.GetCategoryByID(id)
	if err != nil {
		return c.JSON(500, echo.Map{"message": err.Error()})
	}
	return c.JSON(200, category)
}

func (h *CategoryHandler) PatchCategory(c echo.Context, category *request.CategoryPatchRequest) error {
	id := c.Param("id")

	updatedCategory, err := h.categoryUsecase.UpdateCategory(id, category)
	if err != nil {
		switch {
		case errors.Is(err, pkg.CategoryNotFound):
			return c.JSON(404, echo.Map{"message": "Category not found"})
		case errors.Is(err, pkg.ParentCategoryNotFound):
			return c.JSON(404, echo.Map{"message": "Parent category not found"})
		case errors.Is(err, pkg.CategoryCannotBeItsOwnParent):
			return c.JSON(400, echo.Map{"message": "Category cannot be its own parent"})
		case errors.Is(err, pkg.DuplicateEntry):
			return c.JSON(409, echo.Map{"message": "Category name already exists"})
		case errors.Is(err, pkg.NoFieldsToUpdate):
			return c.JSON(400, echo.Map{"message": "No fields provided to update"})
		default:
			return c.JSON(500, echo.Map{"message": "Failed to update category"})
		}
	}

	return c.JSON(200, echo.Map{"message": "Category updated successfully", "data": updatedCategory})
}

func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	if err := h.categoryUsecase.DeleteCategory(id); err != nil {
		return c.JSON(500, echo.Map{"message": err.Error()})
	}
	return c.JSON(204, nil)
}

func (h *CategoryHandler) GetCategoryTree(c echo.Context) error {
	categories, err := h.categoryUsecase.GetCategoryTree()
	if err != nil {
		return c.JSON(500, echo.Map{"message": "Failed to retrieve category tree"})
	}
	return c.JSON(200, echo.Map{"data": categories})
}

func (h *CategoryHandler) GetChildCategoriesByID(c echo.Context) error {
	categories, err := h.categoryUsecase.GetChildCategoriesByID(c.Param("id"))
	if err != nil {
		return c.JSON(500, echo.Map{"message": "Failed to retrieve child categories"})
	}
	return c.JSON(200, echo.Map{"data": categories})
}

func (h *CategoryHandler) GetProductsByCategoryID(c echo.Context) error {
	return nil
}
