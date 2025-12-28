package router

import (
	"product-service/internal/handler"
	"product-service/internal/pkg"
	"product-service/internal/repository"
	"product-service/internal/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterCategoryRoutes(e *echo.Echo, db *gorm.DB) {
	categoryGroup := e.Group("/categories")

	categoryRepo := repository.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	categoryGroup.GET("", categoryHandler.GetAllCategories)
	categoryGroup.GET("/:id", categoryHandler.GetCategoryByID)
	categoryGroup.POST("", pkg.BindAndValidate(categoryHandler.AddCategory))
	categoryGroup.PATCH("/:id", pkg.BindAndValidate(categoryHandler.PatchCategory))
	categoryGroup.DELETE("/:id", categoryHandler.DeleteCategory)
	categoryGroup.GET("/tree", categoryHandler.GetCategoryTree)
	categoryGroup.GET("/leaf", categoryHandler.GetLeafCategories)
	categoryGroup.GET("/:id/children", categoryHandler.GetChildCategoriesByID)

	categoryGroup.GET("/:id/products", categoryHandler.GetProductsByCategoryID)
}

func RegisterProductRoutes(e *echo.Echo, db *gorm.DB) {
	// productGroup := e.Group("/products")
}
