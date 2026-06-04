package admin

import (
	"context"
	"sort"
	"strings"

	"github.com/keto-granola/server/internal/product"
)

type ProductService struct {
	store product.Repository
}

func NewService(store product.Repository) *ProductService {
	return &ProductService{store: store}
}

func (s *ProductService) CreateProduct(ctx context.Context, params *product.CreateProductParams) (*CreateProductResponse, error) {
	createProdParams := &product.CreateProductParams{
		Name:            strings.TrimSpace(params.Name),
		Description:     strings.TrimSpace(params.Description),
		Ingredients:     normaliseIngredients(params.Ingredients),
		Nutrition:       params.Nutrition,
		DietaryTags:     params.DietaryTags,
		WeightG:         params.WeightG,
		Allergens:       params.Allergens,
		PriceCents:      params.PriceCents,
		Currency:        strings.ToUpper(params.Currency),
		ImageStorageKey: params.ImageStorageKey,
		ImageAlt:        strings.TrimSpace(params.ImageAlt),
	}

	newProd, err := s.store.InsertProduct(ctx, createProdParams)
	if err != nil {
		return nil, err
	}

	return &CreateProductResponse{
		ID:          newProd.ID,
		Name:        newProd.Name,
		Description: newProd.Description,
		Ingredients: newProd.Ingredients,
		Nutrition:   newProd.Nutrition,
		WeightG:     newProd.WeightG,
		DietaryTags: newProd.DietaryTags,
		Allergens:   newProd.Allergens,
		PriceCents:  newProd.PriceCents,
		Currency:    newProd.Currency,
		ImageURL:    getCDNUploadURL(newProd.ImageStorageKey),
		ImageAlt:    newProd.ImageAlt,
	}, nil
}

func normaliseIngredients(ingredients []product.Ingredient) []product.Ingredient {
	seenMain := make(map[string]struct{}, len(ingredients))
	out := make([]product.Ingredient, 0, len(ingredients))

	for _, main := range ingredients {
		name := strings.TrimSpace(main.Name)
		if name == "" {
			continue
		}

		name = strings.ToLower(name)

		if _, exists := seenMain[name]; exists {
			continue
		}

		seenMain[name] = struct{}{}

		seenSub := make(map[string]struct{}, len(main.SubIngredients))
		subs := make([]string, 0, len(main.SubIngredients))

		for _, sub := range main.SubIngredients {
			sub = strings.TrimSpace(sub)
			if sub == "" {
				continue
			}

			sub = strings.ToLower(sub)

			if _, exists := seenSub[sub]; exists {
				continue
			}

			seenSub[sub] = struct{}{}

			subs = append(subs, sub)
		}

		out = append(out, product.Ingredient{
			Name:           name,
			SubIngredients: subs,
			Percentage:     main.Percentage,
		})
	}

	sort.SliceStable(out, func(i, j int) bool {
		return out[i].Percentage > out[j].Percentage
	})

	return out
}

// TODO: implement
func getCDNUploadURL(storageKey string) string {
	return storageKey
}
