package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/keto-granola/server/internal/config"
	"github.com/keto-granola/server/internal/server"
	"github.com/keto-granola/server/internal/store/db"
)

func main() {
	if err := run(); err != nil {
		slog.Error("run", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("shutting down gracefully...")
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.ParseEnv()
	if err != nil {
		return fmt.Errorf("parse env vars %v", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.LogLevel,
	}))
	slog.SetDefault(logger)

	db, err := db.New(ctx)
	if err != nil {
		return fmt.Errorf("connect to db %v", err)
	}

	server := server.New(ctx, server.Dependencies{Db: db}, cfg.Environment, cfg.ClientURL)
	svrErr := make(chan error, 1)

	go func() {
		svrErr <- server.Start(cfg.Port)
	}()

	select {
	case <-ctx.Done():
		slog.Info("context cancelled")
	case err = <-svrErr:
	}

	if shutdownErr := server.Stop(); shutdownErr != nil {
		slog.Error("server shutdown", slog.Any("error", shutdownErr))
	}

	return err
}
