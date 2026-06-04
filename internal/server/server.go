package server

import (
	"context"
	"embed"
	"html/template"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/keto-granola/server/internal/config"
	"github.com/keto-granola/server/internal/middleware"
	productadmin "github.com/keto-granola/server/internal/product/admin"
	productweb "github.com/keto-granola/server/internal/product/web"
	"github.com/keto-granola/server/internal/server/templates/templatehelpers"
	"github.com/keto-granola/server/internal/store"
)

const (
	// how long the server will wait to read the entire request after the connection is accepted
	readTimeout = 10 * time.Second
	// how long the server has to write the response after reading the request
	writeTimeout = 10 * time.Second
	// how long to keep a keep-alive connection open waiting for the next request
	idleTimeout     = 120 * time.Second
	shutdownTimeout = 10 * time.Second

	serverRateLimit  = 60
	serverBurstLimit = 120
)

type Server struct {
	echo      *echo.Echo
	templates *template.Template
}

type Handlers struct {
	// api handlers
	ProductAdmin *productadmin.Handler

	// web handlers
	Product *productweb.Handler
}

//go:embed templates
var templateFS embed.FS

func New(ctx context.Context, environment config.Environment, clientURL string, handlers *Handlers, dataStore *store.Store) *Server {
	e := echo.New()
	NewValidator(e)
	e.HideBanner = true // prevents startup banner from being logged

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{clientURL},
	}))

	e.Use(middleware.Log)

	// limits each unique IP to 60 requests per minute with a burst of 120.
	e.Use(echoMiddleware.RateLimiter(echoMiddleware.NewRateLimiterMemoryStoreWithConfig(
		echoMiddleware.RateLimiterMemoryStoreConfig{
			Rate:      serverRateLimit,
			Burst:     serverBurstLimit,
			ExpiresIn: time.Minute,
		},
	)))

	var templates = template.Must(
		template.New("").Funcs(templatehelpers.FuncMap()).ParseFS(templateFS, "internal/templates/**/*.html"),
	)

	web := e.Group("/")
	api := e.Group(config.APIBasePath)

	apiPublic := api.Group("")
	apiPrivate := api.Group("")

	if environment == config.EnvironmentTest {
		// TODO: run test middleware
		slog.Info("run test middleware")
	} else {
		// TODO: run auth middleware
		slog.Info("run auth middleware")
	}

	registerRoutes(apiPublic, apiPrivate, web, handlers, dataStore)

	e.Server.ReadTimeout = readTimeout
	e.Server.WriteTimeout = writeTimeout
	e.Server.IdleTimeout = idleTimeout

	return &Server{
		echo:      e,
		templates: templates,
	}
}

func (s *Server) Start(port string) error {
	if err := s.echo.Start(":" + port); err != nil && err != http.ErrServerClosed {
		slog.Error("start server", slog.Any("error", err))
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	return s.echo.Shutdown(ctx)
}
