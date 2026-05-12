package server

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) AddProduct(e echo.Context) (uuid.UUID, error) {
	return uuid.New(), nil
}
