package product

import (
	"context"

	"github.com/google/uuid"
)

const (
	DietaryTagGlutenFree = "gluten-free"
	DietaryTagKeto       = "keto"
)

//go:generate moq -out mocks/mock.go -pkg mocks . Repository

type Repository interface {
	InsertProduct(ctx context.Context, params *CreateProductParams) (*Product, error)
}

type Product struct {
	ID              uuid.UUID    `json:"id"`
	Name            string       `json:"name"`
	Description     string       `json:"description"`
	Ingredients     []Ingredient `json:"ingredients"`
	Nutrition       Nutrition    `json:"nutrition"`
	WeightG         int32        `json:"weight_g"`
	DietaryTags     []string     `json:"dietary_tags"`
	Allergens       []string     `json:"allergens"`
	PriceCents      int32        `json:"price_cents"`
	Currency        string       `json:"currency"`
	ImageStorageKey string       `json:"image_storage_key"`
	ImageAlt        string       `json:"image_alt"`
}

type Nutrition struct {
	Per100g  NutritionValues `json:"per_100g" validate:"required"`
	ServingG int32           `json:"serving_g" validate:"required"`
}

type NutritionValues struct {
	Calories      float64 `json:"calories" validate:"required"`
	Kilojoules    float64 `json:"kilojoules" validate:"required"`
	ProteinG      float64 `json:"protein_g" validate:"required"`
	FatTotalG     float64 `json:"fat_total_g" validate:"required"`
	FatSaturatedG float64 `json:"fat_saturated_g" validate:"required"`
	CarbsG        float64 `json:"carbs_g" validate:"required"`
	FibreG        float64 `json:"fibre_g" validate:"required"`
	SugarG        float64 `json:"sugar_g" validate:"required"`
	SodiumMg      float64 `json:"sodium_mg" validate:"required"`
}

type Ingredient struct {
	Name           string   `json:"name" validate:"required"`
	SubIngredients []string `json:"sub_ingredients" validate:"required"`
	Percentage     float64  `json:"percentage" validate:"required"`
}

type CreateProductParams struct {
	Name            string       `json:"name" validate:"required"`
	Description     string       `json:"description" validate:"required"`
	Ingredients     []Ingredient `json:"ingredients" validate:"required,min=1,dive"`
	Nutrition       Nutrition    `json:"nutrition" validate:"required"`
	WeightG         int32        `json:"weight_g" validate:"required"`
	DietaryTags     []string     `json:"dietary_tags" validate:"required"`
	Allergens       []string     `json:"allergens" validate:"required"`
	PriceCents      int32        `json:"price_cents" validate:"required"`
	Currency        string       `json:"currency" validate:"required,len=3"`
	ImageStorageKey string       `json:"image_storage_key" validate:"required"`
	ImageAlt        string       `json:"image_alt" validate:"required"`
}
