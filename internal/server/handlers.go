package server

import "github.com/keto-granola/server/internal/admin/product"

type Handlers struct {
	Product product.Repository
}

func NewHandlers(product product.Repository) *Handlers {
	return &Handlers{
		Product: product,
	}
}
