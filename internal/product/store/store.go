package store

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/keto-granola/server/internal/apperr"
	"github.com/keto-granola/server/internal/product"
	"github.com/keto-granola/server/internal/store"
	"github.com/keto-granola/server/internal/store/db/generated"
	"github.com/keto-granola/server/internal/store/db/utils"
)

type Store struct {
	queries *generated.Queries
}

func New(queries *generated.Queries) *Store {
	return &Store{queries: queries}
}

func (s *Store) InsertProduct(ctx context.Context, params *product.CreateProductParams) (*product.Product, error) {
	insertParams, err := toInsertProductParams(params)
	if err != nil {
		return nil, err
	}

	row, err := store.ExecQuery(ctx, func() (generated.InsertProductRow, error) {
		return s.queries.InsertProduct(ctx, *insertParams)
	})

	if err != nil {
		return nil, apperr.Internal("Store.InsertProduct", err)
	}

	return insertedProductFrom(&row)
}

func toInsertProductParams(params *product.CreateProductParams) (*generated.InsertProductParams, error) {
	ingredients, err := json.Marshal(params.Ingredients)
	if err != nil {
		return nil, apperr.Internal("Store.InsertProduct", err)
	}

	nutrition, err := json.Marshal(params.Nutrition)
	if err != nil {
		return nil, apperr.Internal("Store.InsertProduct", err)
	}

	return &generated.InsertProductParams{
		Name:            params.Name,
		Description:     params.Description,
		Ingredients:     ingredients,
		Nutrition:       nutrition,
		WeightG:         params.WeightG,
		DietaryTags:     params.DietaryTags,
		Allergens:       params.Allergens,
		PriceCents:      params.PriceCents,
		Currency:        params.Currency,
		ImageStorageKey: params.ImageStorageKey,
		ImageAlt:        params.ImageAlt,
	}, nil
}

func insertedProductFrom(row *generated.InsertProductRow) (*product.Product, error) {
	var ingredients []product.Ingredient
	if err := json.Unmarshal(row.Ingredients, &ingredients); err != nil {
		return nil, apperr.Internal("Store.InsertProduct", err)
	}

	var nutrition product.Nutrition
	if err := json.Unmarshal(row.Nutrition, &nutrition); err != nil {
		return nil, apperr.Internal("Store.InsertProduct", err)
	}

	return &product.Product{
		ID:              utils.UUIDFrom(row.ID),
		Name:            row.Name,
		Description:     row.Description,
		Ingredients:     ingredients,
		Nutrition:       nutrition,
		WeightG:         row.WeightG,
		DietaryTags:     row.DietaryTags,
		Allergens:       row.Allergens,
		PriceCents:      row.PriceCents,
		Currency:        row.Currency,
		ImageStorageKey: row.ImageStorageKey,
		ImageAlt:        row.ImageAlt,
	}, nil
}

func fetchedProductFrom(row *generated.GetProductRow) (*product.Product, error) {
	var ingredients []product.Ingredient
	if err := json.Unmarshal(row.Ingredients, &ingredients); err != nil {
		return nil, apperr.Internal("Store.GetProduct", err)
	}

	var nutrition product.Nutrition
	if err := json.Unmarshal(row.Nutrition, &nutrition); err != nil {
		return nil, apperr.Internal("Store.GetProduct", err)
	}

	return &product.Product{
		ID:              utils.UUIDFrom(row.ID),
		Name:            row.Name,
		Description:     row.Description,
		Ingredients:     ingredients,
		Nutrition:       nutrition,
		WeightG:         row.WeightG,
		DietaryTags:     row.DietaryTags,
		Allergens:       row.Allergens,
		PriceCents:      row.PriceCents,
		Currency:        row.Currency,
		ImageStorageKey: row.ImageStorageKey,
		ImageAlt:        row.ImageAlt,
	}, nil
}

func (s *Store) GetProduct(ctx context.Context, ID pgtype.UUID) (*product.Product, error) {
	row, err := store.ExecQuery(ctx, func() (generated.GetProductRow, error) {
		return s.queries.GetProduct(ctx, ID)
	})

	if err != nil {
		return nil, apperr.Internal("Store.GetProduct", err)
	}

	return fetchedProductFrom(&row)
}
