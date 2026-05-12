package server

import "github.com/labstack/echo/v4"

func registerRoutes(public, private *echo.Group, handlers *Handlers) {
	// TODO: add healthcheck

	// admin routes
	private.POST("/admin/product", handlers.AddProduct)
}
