package db

import (
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	pgxv5 "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func RunMigrations(dbURL string) error {
	cfg, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return fmt.Errorf("parse db connection config: %w", err)
	}

	db := stdlib.OpenDB(*cfg)
	defer db.Close()

	dbDriver, err := pgxv5.WithInstance(db, &pgxv5.Config{})
	if err != nil {
		return fmt.Errorf("create db driver: %w", err)
	}

	migrationsDriver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("set up migrations driver: %w", err)
	}

	m, err := migrate.NewWithInstance(
		"iofs",
		migrationsDriver,
		"pgx",
		dbDriver,
	)
	if err != nil {
		return fmt.Errorf("create migrate instance: %w", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
