package server

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/keto-granola/server/internal/apperr"
)

type customValidator struct {
	validator *validator.Validate
}

func NewValidator(instance *echo.Echo) {
	instance.Validator = &customValidator{validator: validator.New()}
}

func (cv *customValidator) Validate(i any) error {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		return apperr.Validation("request.Validate", "VALIDATION_ERROR", "invalid request")
	}

	return apperr.Internal("request.Validate", err)
}
