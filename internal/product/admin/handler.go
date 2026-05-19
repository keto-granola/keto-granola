package admin

import (
	"context"

	"github.com/keto-granola/server/internal/product"
)

type Handler struct {
	service *ProductService
}

func NewHandler(s *ProductService) *Handler {
	return &Handler{service: s}
}

type CreateProductRequest struct {
	Name string `json:"name" validate:"required"`
}

func (h *Handler) CreateProduct(ctx context.Context, req CreateProductRequest) (*product.Product, error) {
	prod, err := h.service.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return prod, nil
}
