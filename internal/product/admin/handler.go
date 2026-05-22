package admin

import (
	"context"

	"github.com/google/uuid"

	"github.com/keto-granola/server/internal/product"
)

type Handler struct {
	service *ProductService
}

func NewHandler(s *ProductService) *Handler {
	return &Handler{service: s}
}

type CreateProductResponse struct {
	ID          uuid.UUID            `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Ingredients []product.Ingredient `json:"ingredients"`
	Nutrition   product.Nutrition    `json:"nutrition"`
	WeightG     int32                `json:"weight_g"`
	DietaryTags []string `json:"dietary_tags"`
	Allergens   []string             `json:"allergens"`
	PriceCents  int32                `json:"price_cents"`
	Currency    string               `json:"currency"`
	ImageURL    string               `json:"image_url"`
	ImageAlt    string               `json:"image_alt"`
}

func (h *Handler) CreateProduct(ctx context.Context, params *product.CreateProductParams) (*CreateProductResponse, error) {
	prod, err := h.service.CreateProduct(ctx, params)
	if err != nil {
		return nil, err
	}

	return prod, nil
}
