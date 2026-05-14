package store

import (
	"context"

	"github.com/keto-granola/server/internal/product"
	"github.com/keto-granola/server/internal/store/db/generated"
	"github.com/keto-granola/server/internal/store/db/utils"
)

type Store struct {
	db *generated.Queries
}

func New(db *generated.Queries) *Store {
	return &Store{db: db}
}

func (s *Store) InsertProduct(ctx context.Context, prod *product.Product) (*product.Product, error) {
	ID, err := s.db.InsertProduct(ctx, utils.PGTextFrom(prod.Name))

	if err != nil {
		return nil, err
	}

	prod.ID = utils.UUIDFrom(ID)
	return prod, nil
}
