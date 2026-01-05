package server

import (
	"product-service/config"
	router "product-service/internal"
	"product-service/internal/entity"
	"product-service/internal/middleware"

	"myshop-shared/pkg"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func StartServer() {
	db := config.ConnectDB()

	db.AutoMigrate(
		&entity.Category{},
		&entity.Product{},
		&entity.Attribute{},
		&entity.AttributeValue{},
		&entity.Variant{},
		&entity.ProductImage{},
	)

	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}

	middleware.RegisterBasicMiddleware(e)

	router.RegisterHealthCheckRoute(e)
	router.RegisterCategoryRoutes(e, db)
	router.RegisterProductRoutes(e, db)
	router.RegisterAttributeRoutes(e, db)
	router.RegisterAttributeValueRoutes(e, db)
	router.RegisterVariantRoutes(e, db)
	router.RegisterProductImageRoutes(e, db)

	e.Logger.Fatal(e.Start(":8081"))
}
