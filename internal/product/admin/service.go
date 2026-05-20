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

func (s *ProductService) CreateProduct(ctx context.Context, req *CreateProductRequest) (*CreateProductResponse, error) {
	prod := &product.Product{
		Name:            strings.TrimSpace(req.Name),
		Description:     strings.TrimSpace(req.Description),
		Ingredients:     normaliseIngredients(req.Ingredients),
		Nutrition:       req.Nutrition,
		DietaryTags:     normaliseDietaryTags(req.DietaryTags),
		WeightG:         req.WeightG,
		Allergens:       normaliseAllergens(req.Allergens),
		PriceCents:      req.PriceCents,
		Currency:        strings.ToUpper(req.Currency),
		ImageStorageKey: req.ImageStorageKey,
		ImageAlt:        strings.TrimSpace(req.ImageAlt),
	}

	newProd, err := s.store.InsertProduct(ctx, prod)
	if err != nil {
		return nil, err
	}

	// TODO: construct cloudFront URL
	cdnURL := "" // e.g: d3f8k9x7abc123.cloudfront.net/images/product-123/image-1.jpg

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
		ImageURL:    cdnURL,
		ImageAlt:    newProd.ImageAlt,
	}, nil
}

func normaliseIngredients(ingredients []product.Ingredient) []product.Ingredient {
	seen := make(map[string]struct{}, len(ingredients))
	out := make([]product.Ingredient, 0, len(ingredients))

	for _, ing := range ingredients {
		name := strings.TrimSpace(ing.Name)
		if name == "" {
			continue
		}

		name = strings.ToLower(name)
		if _, exists := seen[name]; exists {
			continue
		}

		seen[name] = struct{}{}

		subIngredients := make([]string, 0, len(ing.SubIngredients))
		for _, sub := range ing.SubIngredients {
			sub = strings.TrimSpace(sub)
			if sub == "" {
				continue
			}
			subIngredients = append(subIngredients, sub)
		}

		out = append(out, product.Ingredient{
			Name:           name,
			SubIngredients: subIngredients,
			Percentage:     ing.Percentage,
		})
	}

	sort.SliceStable(out, func(i, j int) bool {
		return out[i].Percentage > out[j].Percentage
	})

	return out
}

func normaliseDietaryTags(dietaryTags []product.DietaryTag) []product.DietaryTag {
	seen := make(map[string]struct{}, len(dietaryTags))
	out := make([]product.DietaryTag, 0, len(dietaryTags))

	for _, tag := range dietaryTags {
		name := strings.TrimSpace(string(tag))

		if name == "" {
			continue
		}

		name = strings.ToUpper(name)
		if _, exists := seen[name]; exists {
			continue
		}

		seen[name] = struct{}{}
		out = append(out, product.DietaryTag(name))
	}

	return out
}

func normaliseAllergens(allergens []string) []string {
	seen := make(map[string]struct{}, len(allergens))
	out := make([]string, 0, len(allergens))

	for _, allergen := range allergens {
		name := strings.TrimSpace(allergen)

		if name == "" {
			continue
		}

		name = strings.ToLower(name)
		if _, exists := seen[name]; exists {
			continue
		}

		seen[name] = struct{}{}
		out = append(out, name)
	}

	return out
}
