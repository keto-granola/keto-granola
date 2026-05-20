package product

import (
	"context"
	"time"

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
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Description string        `json:"description"`
	Ingredients []Ingredient  `json:"ingredients"`
	Nutrition   NutritionInfo `json:"nutrition"`
	DietaryTags []DietaryTag  `json:"dietary_tags"`
	Allergens   []string      `json:"allergens"`
	PriceCents  int64         `json:"price_cents"`
	Currency    string        `json:"currency"`
	Image_URL   string        `json:"image_url"`
	Image_ALT   string        `json:"image_alt"`
}

type NutritionInfo struct {
	ServingSize string `json:"serving_size"`

	Calories      float64 `json:"calories"`
	Kilojoules    float64 `json:"kilojoules"`
	ProteinG      float64 `json:"protein_g"`
	FatTotalG     float64 `json:"fat_total_g"`
	FatSaturatedG float64 `json:"fat_saturated_g"`
	CarbsG        float64 `json:"carbs_g"`
	FiberG        float64 `json:"fiber_g"`
	SugarG        float64 `json:"sugar_g"`
	SodiumMg      float64 `json:"sodium_mg"`
}

type Ingredient struct {
	Name       string   `json:"name"`
	Percentage *float64 `json:"percentage,omitempty"`
}
