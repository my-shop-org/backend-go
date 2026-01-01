package server

import (
	"product-service/config"
	router "product-service/internal"
	"product-service/internal/entity"
	"product-service/internal/middleware"
	"product-service/internal/pkg"

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
	)

	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}

	middleware.RegisterBasicMiddleware(e)

	router.RegisterCategoryRoutes(e, db)
	router.RegisterProductRoutes(e, db)
	router.RegisterAttributeRoutes(e, db)
	router.RegisterAttributeValueRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
