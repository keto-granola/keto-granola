package product

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const (
	DietaryTagGlutenFree = "gluten-free"
	DietaryTagKeto       = "keto"
)

//go:generate moq -out mocks/mock.go -pkg mocks . Repository

type Repository interface {
	GetProduct(ctx context.Context, ID pgtype.UUID) (*Product, error)
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
	SubIngredients []string `json:"sub_ingredients"`
	Percentage     float64  `json:"percentage" validate:"required"`
}
