package pkg

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

type ValidationErrors struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func BindAndValidate[T any](fn func(echo.Context, *T) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		rawRequest := new(T)

		if err := c.Bind(rawRequest); err != nil {
			return c.JSON(400, echo.Map{"error": "Invalid request"})
		}

		if err := c.Validate(rawRequest); err != nil {
			validationErrors := FormatValidationError(err)
			return c.JSON(400, echo.Map{"errors": validationErrors})
		}

		return fn(c, rawRequest)
	}
}

func FormatValidationError(err error) []ValidationErrors {
	var ve validator.ValidationErrors
	if ok := errors.As(err, &ve); !ok {
		return []ValidationErrors{{Message: err.Error()}}
	}

	errArr := make([]ValidationErrors, 0, len(ve))
	for _, fe := range ve {
		errArr = append(errArr, ValidationErrors{
			Field:   fe.Field(),
			Message: GetErrorMessage(fe),
		})
	}
	return errArr
}

func GetErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Value must be at least %s characters long", fe.Param())
	case "max":
		return fmt.Sprintf("Value must be at most %s characters long", fe.Param())
	}
	return fe.Error()
}
