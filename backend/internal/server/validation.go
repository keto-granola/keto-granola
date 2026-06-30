package server

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/keto-granola/keto-granola/internal/apperr"
)

type customValidator struct {
	validator *validator.Validate
}

func NewValidator(e *echo.Echo) {
	e.Validator = &customValidator{validator: validator.New()}
}

func (cv *customValidator) Validate(i any) error {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		return apperr.Validation("request.Validate", "VALIDATION_ERROR", apperr.ErrMsgValidation)
	}

	return apperr.Internal("request.Validate", err)
}
