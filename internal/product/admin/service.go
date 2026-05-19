package admin

import (
	"context"

	"github.com/keto-granola/server/internal/product"
)

type ProductService struct {
	store product.Repository
}

func NewService(store product.Repository) *ProductService {
	return &ProductService{store: store}
}

func (s *ProductService) CreateProduct(ctx context.Context, req CreateProductRequest) (*product.Product, error) {
	p := &product.Product{
		Name: req.Name,
	}

	return s.store.InsertProduct(ctx, p)
}
