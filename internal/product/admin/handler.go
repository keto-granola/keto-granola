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
	Name            string                `json:"name" validate:"required"`
	Description     string                `json:"description" validate:"required"`
	Ingredients     []product.Ingredient  `json:"ingredients" validate:"required,min=1,dive"`
	NutritionalInfo product.NutritionInfo `json:"nutritional_info" validate:"required"`
	DietaryTags     []product.DietaryTag  `json:"dietary_tags,omitempty" validate:"dive,dietary_tag"`
	Allergens       []string              `json:"allergens,omitempty"`
	PriceCents      int64                 `json:"price_cents" validate:"required"`
	Currency        string                `json:"currency" validate:"required,len=3"`
	Image_ALT       string                `json:"image_alt" validate:"required"`
}

func (h *Handler) CreateProduct(ctx context.Context, req CreateProductRequest) (*product.Product, error) {
	prod, err := h.service.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return prod, nil
}
