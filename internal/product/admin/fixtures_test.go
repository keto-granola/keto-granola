package admin_test

import (
	"github.com/google/uuid"
	"github.com/keto-granola/server/internal/product"
	"github.com/keto-granola/server/internal/product/admin"
)

var createProductReqBody = &product.CreateProductParams{
	Name:        "Test Granola  ",
	Description: "   Test Description",
	Ingredients: []product.Ingredient{
		{Name: "iNG1", SubIngredients: []string{"sub1", "sub2", "sub3", "sub4"}, Percentage: 70},
		{Name: "ing1", SubIngredients: []string{}, Percentage: 70},
		{Name: "ing3 ", SubIngredients: []string{"sub1", "sub2", "SUB1", "sub3 ", "  suB4", "Sub1"}, Percentage: 10},
		{Name: "ING2", SubIngredients: []string{}, Percentage: 20},
	},
	Nutrition: product.Nutrition{
		Per100g: product.NutritionValues{
			Kilojoules:    800,
			Calories:      400,
			FatTotalG:     30,
			FatSaturatedG: 15,
			ProteinG:      10,
			CarbsG:        18,
			FibreG:        10,
			SugarG:        8,
			SodiumMg:      125,
		},
		ServingG: 40,
	},
	WeightG:         450,
	DietaryTags:     []string{"keto", "gluten-free"},
	Allergens:       []string{},
	PriceCents:      1200,
	Currency:        "aud",
	ImageStorageKey: "test/123hsbd-key",
	ImageAlt:        "test alt  ",
}

var expectedCreateProductParams = &product.CreateProductParams{
	Name:        "Test Granola",
	Description: "Test Description",
	Ingredients: []product.Ingredient{
		{Name: "ing1", SubIngredients: []string{"sub1", "sub2", "sub3", "sub4"}, Percentage: 70},
		{Name: "ing2", SubIngredients: []string{}, Percentage: 20},
		{Name: "ing3", SubIngredients: []string{"sub1", "sub2", "sub3", "sub4"}, Percentage: 10},
	},
	Nutrition: product.Nutrition{
		Per100g: product.NutritionValues{
			Kilojoules:    800,
			Calories:      400,
			FatTotalG:     30,
			FatSaturatedG: 15,
			ProteinG:      10,
			CarbsG:        18,
			FibreG:        10,
			SugarG:        8,
			SodiumMg:      125,
		},
		ServingG: 40,
	},
	WeightG:         450,
	DietaryTags:     []string{"keto", "gluten-free"},
	Allergens:       []string{},
	PriceCents:      1200,
	Currency:        "AUD",
	ImageStorageKey: "test/123hsbd-key",
	ImageAlt:        "test alt",
}

var insertedProd = &product.Product{
	ID:          uuid.New(),
	Name:        "Test Granola",
	Description: "Test Description",
	Ingredients: []product.Ingredient{
		{Name: "ing1", SubIngredients: []string{"sub1", "sub2", "sub3", "sub4"}, Percentage: 70},
		{Name: "ing2", SubIngredients: []string{}, Percentage: 20},
		{Name: "ing3", SubIngredients: []string{"sub1", "sub2", "sub3", "sub4"}, Percentage: 10},
	},
	Nutrition: product.Nutrition{
		Per100g: product.NutritionValues{
			Kilojoules:    800,
			Calories:      400,
			FatTotalG:     30,
			FatSaturatedG: 15,
			ProteinG:      10,
			CarbsG:        18,
			FibreG:        10,
			SugarG:        8,
			SodiumMg:      125,
		},
		ServingG: 40,
	},
	DietaryTags:     []string{"keto", "gluten-free"},
	Allergens:       []string{},
	WeightG:         450,
	PriceCents:      1200,
	Currency:        "AUD",
	ImageAlt:        "test alt",
	ImageStorageKey: "test/183839-key",
}

var expectedCreateProducRes = admin.CreateProductResponse{
	ID:          insertedProd.ID,
	Name:        insertedProd.Name,
	Description: insertedProd.Description,
	Ingredients: insertedProd.Ingredients,
	Nutrition:   insertedProd.Nutrition,
	DietaryTags: insertedProd.DietaryTags,
	Allergens:   insertedProd.Allergens,
	WeightG:     insertedProd.WeightG,
	PriceCents:  insertedProd.PriceCents,
	Currency:    insertedProd.Currency,
	ImageAlt:    insertedProd.ImageAlt,
	ImageURL:    "",
}
