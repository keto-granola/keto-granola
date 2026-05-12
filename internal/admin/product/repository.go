package product

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Add(ctx context.Context, product *Product) (ID uuid.UUID, err error)
}
