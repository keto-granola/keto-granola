package main

import (
	"context"
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/keto-granola/server/internal/config"
	"github.com/keto-granola/server/internal/product"
	productadmin "github.com/keto-granola/server/internal/product/admin"
	productstore "github.com/keto-granola/server/internal/product/store"
	"github.com/keto-granola/server/internal/product/web"
	"github.com/keto-granola/server/internal/server"
	"github.com/keto-granola/server/internal/store"
)

func main() {
	if err := run(); err != nil {
		slog.Error("run", slog.Any("error", err))
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.ParseEnv()
	if err != nil {
		return fmt.Errorf("parse env vars %w", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.LogLevel,
	}))
	slog.SetDefault(logger)

	dataStore, err := store.New(ctx, cfg.DbURL)
	if err != nil {
		return fmt.Errorf("create store %w", err)
	}
	defer dataStore.Close()

	templates, err := server.NewTemplates()
	if err != nil {
		return fmt.Errorf("create templates: %w", err)
	}

	handlers := composeHandlers(dataStore, templates, cfg.Environment, cfg.ClientURL)

	serverDeps := &server.Deps{
		Environment: cfg.Environment,
		ClientURL:   cfg.ClientURL,
		Handlers:    handlers,
		DataStore:   dataStore,
	}

	echo := server.New(ctx, serverDeps)

	serverErr := make(chan error, 1)

	go func() {
		serverErr <- echo.Start(cfg.Port)
	}()

	select {
	case <-ctx.Done():
		slog.Info("context cancelled")
	case err = <-serverErr:
	}

	if shutdownErr := echo.Stop(); shutdownErr != nil {
		slog.Error("server shutdown", slog.Any("error", shutdownErr))
	}

	return err
}

func composeHandlers(db *store.Store, tmpl *template.Template, env config.Environment, clientURL string) *server.Handlers {
	productStore := productstore.New(db.Queries)
	prodService := product.NewService(productStore)
	prodAdminService := productadmin.NewService(productStore)

	return &server.Handlers{
		ProductAdmin: productadmin.NewHandler(prodAdminService),
		Product:      web.NewHandler(prodService, tmpl, env, clientURL),
	}
}
