package server

import (
	"product-service/config"
	"product-service/internal"
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
	)

	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}

	middleware.RegisterBasicMiddleware(e)

	internal.RegisterCategoryRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
