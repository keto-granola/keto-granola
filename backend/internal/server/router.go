package server

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/keto-granola/server/internal/store"
)

const pingTimeout = 5 * time.Second

func registerRoutes(apiPublic, apiPrivate, web *echo.Group, handlers *Handlers, dataStore *store.Store) {
	registerHealthEndpoint(apiPublic, dataStore)
	registerAPIRoutes(apiPrivate, handlers)
	registerWebRoutes(web, handlers)
}

func registerHealthEndpoint(api *echo.Group, dataStore *store.Store) {
	api.GET("/health", func(e echo.Context) error {
		dbStatus := "ok"
		httpStatus := http.StatusOK

		pingCtx, cancel := context.WithTimeout(e.Request().Context(), pingTimeout)
		defer cancel()

		err := dataStore.PingDB(pingCtx)
		if err != nil {
			httpStatus = http.StatusServiceUnavailable
			dbStatus = "unreachable"
		}

		return e.JSON(httpStatus, map[string]string{
			"status": "ok",
			"db":     dbStatus,
		})
	})
}

func registerAPIRoutes(apiPrivate *echo.Group, handlers *Handlers) {
	// admin routes
	apiPrivate.POST("/admin/products", Handle(handlers.ProductAdmin.CreateProduct, http.StatusCreated))
}

func registerWebRoutes(web *echo.Group, handlers *Handlers) {
	web.GET("products/:id", handlers.Product.GetProductPage)
}
