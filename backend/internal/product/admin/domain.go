package admin

import (
	"context"

	"github.com/keto-granola/server/internal/product"
)

//go:generate moq -out mocks/mock.go -pkg mocks . Repository

type Repository interface {
	InsertProduct(ctx context.Context, params *CreateProductParams) (*product.Product, error)
}

type CreateProductParams struct {
	Name            string               `json:"name" validate:"required"`
	Description     string               `json:"description" validate:"required"`
	Ingredients     []product.Ingredient `json:"ingredients" validate:"required,min=1,dive"`
	Nutrition       product.Nutrition    `json:"nutrition" validate:"required"`
	WeightG         int32                `json:"weight_g" validate:"required"`
	DietaryTags     []string             `json:"dietary_tags" validate:"required"`
	Allergens       []string             `json:"allergens" validate:"required"`
	PriceCents      int32                `json:"price_cents" validate:"required"`
	Currency        string               `json:"currency" validate:"required,len=3"`
	ImageStorageKey string               `json:"image_storage_key" validate:"required"`
	ImageAlt        string               `json:"image_alt" validate:"required"`
}
