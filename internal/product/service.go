package product

import (
	"context"

	"github.com/google/uuid"
	"github.com/keto-granola/server/internal/store/db/utils"
)

type Service struct {
	store Repository
}

func NewService(store Repository) *Service {
	return &Service{store: store}
}

type NutritionView struct {
	Per100g    NutritionValues
	PerServing NutritionValues
	ServingG   int32
}

type GetProductResponse struct {
	ID          uuid.UUID
	Name        string
	Description string
	Ingredients []Ingredient
	Nutrition   NutritionView
	DietaryTags []string
	Allergens   []string
	Price       float32
	Currency    string
	ImageURL    string
	ImageAlt    string
}

func (s *Service) GetProduct(ctx context.Context, ID uuid.UUID) (*GetProductResponse, error) {
	pgUUID := utils.PGUUIDFromUUID(ID)

	prod, err := s.store.GetProduct(ctx, pgUUID)
	if err != nil {
		return nil, err
	}

	return &GetProductResponse{
		ID:          prod.ID,
		Name:        prod.Name,
		Description: prod.Description,
		Ingredients: prod.Ingredients,
		Nutrition:   toNutritionView(prod.Nutrition),
		DietaryTags: prod.DietaryTags,
		Allergens:   prod.Allergens,
		Price:       float32(prod.PriceCents) / 100,
		Currency:    prod.Currency,
		ImageURL:    getCDNDownloadURL(prod.ImageStorageKey),
		ImageAlt:    prod.ImageAlt,
	}, nil
}

// TODO: implement
func getCDNDownloadURL(storageKey string) string {
	return storageKey
}

func toNutritionView(nutrition Nutrition) NutritionView {
	return NutritionView{
		Per100g: nutrition.Per100g,
		// TODO: implement this
		PerServing: nutrition.Per100g,
		ServingG:   nutrition.ServingG,
	}
}
