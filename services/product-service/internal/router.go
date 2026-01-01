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
}

func RegisterProductRoutes(e *echo.Echo, db *gorm.DB) {
	productGroup := e.Group("/products")

	productRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	productGroup.GET("", productHandler.GetAllProducts)
	productGroup.GET("/:id", productHandler.GetProductByID)
	productGroup.POST("", pkg.BindAndValidate(productHandler.AddProduct))
	productGroup.PATCH("/:id", pkg.BindAndValidate(productHandler.PatchProduct))
	productGroup.DELETE("/:id", productHandler.DeleteProduct)
}

func RegisterAttributeRoutes(e *echo.Echo, db *gorm.DB) {
	attributeGroup := e.Group("/attributes")

	attributeRepo := repository.NewAttributeRepository(db)
	attributeUsecase := usecase.NewAttributeUsecase(attributeRepo)
	attributeHandler := handler.NewAttributeHandler(attributeUsecase)

	attributeGroup.GET("", attributeHandler.GetAllAttributes)
	attributeGroup.GET("/:id", attributeHandler.GetAttributeByID)
	attributeGroup.POST("", pkg.BindAndValidate(attributeHandler.AddAttribute))
	attributeGroup.PATCH("/:id", pkg.BindAndValidate(attributeHandler.PatchAttribute))
	attributeGroup.DELETE("/:id", attributeHandler.DeleteAttribute)
}

func RegisterAttributeValueRoutes(e *echo.Echo, db *gorm.DB) {
	attributeValueGroup := e.Group("/attribute-values")

	attributeValueRepo := repository.NewAttributeValueRepository(db)
	attributeValueUsecase := usecase.NewAttributeValueUsecase(attributeValueRepo)
	attributeValueHandler := handler.NewAttributeValueHandler(attributeValueUsecase)

	attributeValueGroup.GET("", attributeValueHandler.GetAllAttributeValues)
	attributeValueGroup.GET("/:id", attributeValueHandler.GetAttributeValueByID)
	attributeValueGroup.POST("", pkg.BindAndValidate(attributeValueHandler.AddAttributeValue))
	attributeValueGroup.PATCH("/:id", pkg.BindAndValidate(attributeValueHandler.PatchAttributeValue))
	attributeValueGroup.DELETE("/:id", attributeValueHandler.DeleteAttributeValue)
}
