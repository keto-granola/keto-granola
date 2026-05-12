package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	pool    *pgxpool.Pool
	// TODO: implement sqlc
	// Queries *sqlc.Queries
}
