package admin_test

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/keto-granola/keto-granola/internal/apperr"
	"github.com/keto-granola/keto-granola/internal/product"
	"github.com/keto-granola/keto-granola/internal/product/admin"
	"github.com/keto-granola/keto-granola/internal/product/admin/mocks"
)

const (
	ingredient1    = "ing1"
	subIngedient1  = "sub1"
	subIngredient2 = "sub2"
	subIngredient3 = "sub3"
	subIngredient4 = "sub4"
)

var createProdReqBody = &admin.CreateProductParams{
	Name:        "Test Granola  ",
	Description: "   Test Description",
	Ingredients: []product.Ingredient{
		{Name: "iNG1", SubIngredients: []string{subIngedient1, subIngredient2, subIngredient3, subIngredient4}, Percentage: 70},
		{Name: ingredient1, SubIngredients: []string{}, Percentage: 70},
		{Name: "ing3 ", SubIngredients: []string{subIngedient1, subIngredient2, "SUB1", "sub3 ", "  suB4", "Sub1"}, Percentage: 10},
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
	DietaryTags:     []string{product.DietaryTagKeto, product.DietaryTagGlutenFree},
	Allergens:       []string{},
	PriceCents:      1200,
	Currency:        "aud",
	ImageStorageKey: "test/123hsbd-key",
	ImageAlt:        "test alt  ",
}

var expCreateProdParams = &admin.CreateProductParams{
	Name:        "Test Granola",
	Description: "Test Description",
	Ingredients: []product.Ingredient{
		{Name: ingredient1, SubIngredients: []string{subIngedient1, subIngredient2, subIngredient3, subIngredient4}, Percentage: 70},
		{Name: "ing2", SubIngredients: []string{}, Percentage: 20},
		{Name: "ing3", SubIngredients: []string{subIngedient1, subIngredient2, subIngredient3, subIngredient4}, Percentage: 10},
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
	DietaryTags:     []string{product.DietaryTagKeto, product.DietaryTagGlutenFree},
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
		{Name: ingredient1, SubIngredients: []string{subIngedient1, subIngredient2, subIngredient3, subIngredient4}, Percentage: 70},
		{Name: "ing2", SubIngredients: []string{}, Percentage: 20},
		{Name: "ing3", SubIngredients: []string{subIngedient1, subIngredient2, subIngredient3, subIngredient4}, Percentage: 10},
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
	DietaryTags:     []string{product.DietaryTagKeto, product.DietaryTagGlutenFree},
	Allergens:       []string{},
	WeightG:         450,
	PriceCents:      1200,
	Currency:        "AUD",
	ImageAlt:        "test alt",
	ImageStorageKey: "test/183839-key",
}

var expCreateProdRes = admin.CreateProductResponse{
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

type testCase struct {
	name           string
	reqBody        admin.CreateProductParams
	arrange        func() *admin.Handler
	wantHTTPStatus int
	expectedErrMsg string
}

var arrangeValidationCases = func() *admin.Handler {
	mockRepo := &mocks.RepositoryMock{
		InsertProductFunc: func(ctx context.Context, params *admin.CreateProductParams) (*product.Product, error) {
			return insertedProd, nil
		},
	}

	h := admin.NewHandler(admin.NewService(mockRepo))
	return h
}

var createProdUnhappyPathTestCases = []testCase{
	{
		name: "validation - missing ingredient percentage",
		reqBody: admin.CreateProductParams{
			Name:        createProdReqBody.Name,
			Description: createProdReqBody.Description,
			Ingredients: []product.Ingredient{
				{Name: ingredient1, SubIngredients: []string{}, Percentage: 80},
				{Name: "ing2 ", SubIngredients: []string{subIngedient1, subIngredient2, "SUB1", "sub2 ", "Sub1"}},
			},
			Nutrition:       createProdReqBody.Nutrition,
			WeightG:         createProdReqBody.WeightG,
			DietaryTags:     createProdReqBody.DietaryTags,
			Allergens:       createProdReqBody.Allergens,
			PriceCents:      createProdReqBody.PriceCents,
			Currency:        createProdReqBody.Currency,
			ImageStorageKey: createProdReqBody.ImageStorageKey,
			ImageAlt:        createProdReqBody.ImageAlt,
		},
		arrange:        arrangeValidationCases,
		wantHTTPStatus: http.StatusBadRequest,
		expectedErrMsg: apperr.ErrMsgValidation,
	},
	{
		name: "validation - missing ingredients",
		reqBody: admin.CreateProductParams{
			Name:            createProdReqBody.Name,
			Description:     createProdReqBody.Description,
			Ingredients:     []product.Ingredient{},
			Nutrition:       createProdReqBody.Nutrition,
			WeightG:         createProdReqBody.WeightG,
			DietaryTags:     createProdReqBody.DietaryTags,
			Allergens:       createProdReqBody.Allergens,
			PriceCents:      createProdReqBody.PriceCents,
			Currency:        createProdReqBody.Currency,
			ImageStorageKey: createProdReqBody.ImageStorageKey,
			ImageAlt:        createProdReqBody.ImageAlt,
		},
		arrange:        arrangeValidationCases,
		wantHTTPStatus: http.StatusBadRequest,
		expectedErrMsg: apperr.ErrMsgValidation,
	},
	{
		name: "validation - invalid currency",
		reqBody: admin.CreateProductParams{
			Name:            createProdReqBody.Name,
			Description:     createProdReqBody.Description,
			Ingredients:     createProdReqBody.Ingredients,
			Nutrition:       createProdReqBody.Nutrition,
			WeightG:         createProdReqBody.WeightG,
			DietaryTags:     createProdReqBody.DietaryTags,
			Allergens:       createProdReqBody.Allergens,
			PriceCents:      createProdReqBody.PriceCents,
			Currency:        "gb",
			ImageStorageKey: createProdReqBody.ImageStorageKey,
			ImageAlt:        createProdReqBody.ImageAlt,
		},
		arrange:        arrangeValidationCases,
		wantHTTPStatus: http.StatusBadRequest,
		expectedErrMsg: apperr.ErrMsgValidation,
	},
	{
		name:    "internal server error",
		reqBody: *createProdReqBody,
		arrange: func() *admin.Handler {
			mockRepo := &mocks.RepositoryMock{
				InsertProductFunc: func(ctx context.Context, params *admin.CreateProductParams) (*product.Product, error) {
					return nil, errors.New("db internal server error")
				},
			}

			h := admin.NewHandler(admin.NewService(mockRepo))
			return h
		},
		wantHTTPStatus: http.StatusInternalServerError,
		expectedErrMsg: apperr.ErrMsgInternal,
	},
}
