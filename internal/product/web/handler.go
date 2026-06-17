package web

import (
	"html/template"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/keto-granola/server/internal/apperr"
	"github.com/keto-granola/server/internal/config"
	"github.com/keto-granola/server/internal/product"
)

type Handler struct {
	service     *product.Service
	templates   *template.Template
	environment config.Environment
	clientURL   string
}

type ProductPageData struct {
	Product     *product.GetProductResponse
	Environment config.Environment
	ClientURL   string
}

func NewHandler(svc *product.Service, tmpl *template.Template, env config.Environment, clientURL string) *Handler {
	return &Handler{service: svc, templates: tmpl, environment: env, clientURL: clientURL}
}

func (h *Handler) GetProductPage(e echo.Context) error {
	ID, err := uuid.Parse(e.Param("id"))
	if err != nil {
		return apperr.ToHTTPError(apperr.Validation("request.Validate", "VALIDATION_ERROR", apperr.ErrMsgValidation))
	}

	prod, err := h.service.GetProduct(e.Request().Context(), ID)
	if err != nil {
		return apperr.ToHTTPError(err)
	}

	productPageData := &ProductPageData{
		Product:     prod,
		Environment: h.environment,
		ClientURL:   h.clientURL,
	}

	return h.templates.ExecuteTemplate(e.Response(), "product.html", productPageData)
}
