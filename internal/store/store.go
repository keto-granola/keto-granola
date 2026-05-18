package store

import (
	"context"
	"errors"
	"log/slog"
	"slices"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/keto-granola/server/internal/store/db/generated"
	"github.com/keto-granola/server/internal/utils"
)

const (
	DbMaxRetries = 5
	DbBaseDelay  = 100 * time.Millisecond
)

type Store struct {
	pool    *pgxpool.Pool
	Queries *generated.Queries
}

func New(ctx context.Context) (*Store, error) {
	return &Store{
		pool:    nil,
		Queries: nil,
	}, nil
}

func (s *Store) Close() {
	slog.Info("closing pool", slog.Any("pool", s.pool))
}

func ExecQuery[T any](ctx context.Context, query func() (T, error)) (T, error) {
	return utils.RetryWithExponentialBackoff(ctx, query, DbMaxRetries, DbBaseDelay, isRetryableDbError)
}

func ExecCommand(ctx context.Context, command func() error) error {
	_, err := ExecQuery(ctx, func() (*struct{}, error) { return nil, command() })
	return err
}

var transientPostgresErrorCodes = []string{
	"08", // Connection exceptions (network problems, can't reach database)
	"40", // Transaction rollback (like deadlocks or serialisation failures)
	"53", // Insufficient resources (out of memory, disk full)
	"55", // Object not in prerequisite state (like trying to use a prepared statement that doesn't exist)
	"57", // Operator intervention (admin killed the query, database shutting down)
}

func isRetryableDbError(err error) bool {
	if err == nil {
		return false
	}

	if isNotFoundErr(err) {
		return false
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		errClass := pgErr.Code[:2]
		return slices.Contains(transientPostgresErrorCodes, errClass)
	}

	return false
}

func isNotFoundErr(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
