package product

import (
	"context"

	"github.com/google/uuid"
)

type DietaryTag string

const (
	DietaryTagKeto       DietaryTag = "keto"
	DietaryTagGlutenFree DietaryTag = "gluten_free"
	DietaryTagVegetarian DietaryTag = "vegetarian"
)

type Repository interface {
	InsertProduct(ctx context.Context, product *Product) (*Product, error)
}

type Product struct {
	ID              uuid.UUID    `json:"id"`
	Name            string       `json:"name"`
	Description     string       `json:"description"`
	Ingredients     []Ingredient `json:"ingredients"`
	Nutrition       Nutrition    `json:"nutrition"`
	WeightG         int64        `json:"weight_g"`
	DietaryTags     []DietaryTag `json:"dietary_tags"`
	Allergens       []string     `json:"allergens"`
	PriceCents      int64        `json:"price_cents"`
	Currency        string       `json:"currency"`
	ImageStorageKey string       `json:"image_storage_key"`
	ImageAlt        string       `json:"image_alt"`
}

type Nutrition struct {
	Per100g  NutritionValues `json:"per_100g" validate:"required"`
	ServingG int64           `json:"serving_g" validate:"required"`
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
	SubIngredients []string `json:"sub_ingredients,omitempty"`
	Percentage     float64  `json:"percentage"`
}
