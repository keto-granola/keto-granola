package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	pool *pgxpool.Pool
	// TODO: implement sqlc
	// Queries *sqlc.Queries
}

func New(ctx context.Context) (*Db, error) {
	return nil, nil
}
