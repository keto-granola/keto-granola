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

type CreateProductRequest struct {
	Name            string               `json:"name" validate:"required"`
	Description     string               `json:"description" validate:"required"`
	Ingredients     []product.Ingredient `json:"ingredients" validate:"required,min=1,dive"`
	Nutrition       product.Nutrition    `json:"nutrition" validate:"required"`
	WeightG         int64                `json:"weight_g" validate:"required"`
	DietaryTags     []product.DietaryTag `json:"dietary_tags,omitempty" validate:"dive,dietary_tag"`
	Allergens       []string             `json:"allergens,omitempty"`
	PriceCents      int64                `json:"price_cents" validate:"required"`
	Currency        string               `json:"currency" validate:"required,len=3"`
	ImageStorageKey string               `json:"image_storage_key" validate:"required"`
	ImageAlt        string               `json:"image_alt" validate:"required"`
}

type CreateProductResponse struct {
	ID          uuid.UUID            `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Ingredients []product.Ingredient `json:"ingredients"`
	Nutrition   product.Nutrition    `json:"nutrition"`
	WeightG     int64                `json:"weight_g"`
	DietaryTags []product.DietaryTag `json:"dietary_tags"`
	Allergens   []string             `json:"allergens"`
	PriceCents  int64                `json:"price_cents"`
	Currency    string               `json:"currency"`
	ImageURL    string               `json:"image_url"`
	ImageAlt    string               `json:"image_alt"`
}

func (h *Handler) CreateProduct(ctx context.Context, req *CreateProductRequest) (*CreateProductResponse, error) {
	prod, err := h.service.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return prod, nil
}
