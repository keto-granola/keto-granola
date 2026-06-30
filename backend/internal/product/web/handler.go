package web

import (
	"html/template"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/keto-granola/keto-granola/internal/apperr"
	"github.com/keto-granola/keto-granola/internal/config"
	"github.com/keto-granola/keto-granola/internal/product"
)

type Handler struct {
	service   *product.Service
	templates *template.Template
	clientURL string
	devEnv    bool
}

type ProductData struct {
	Product   *product.GetProductResponse
	ClientURL string
	DevEnv    bool
}

func NewHandler(svc *product.Service, tmpl *template.Template, clientURL string, env config.Environment) *Handler {
	return &Handler{service: svc, templates: tmpl, clientURL: clientURL, devEnv: env == config.EnvironmentDevelopment}
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

	productData := &ProductData{
		Product:   prod,
		ClientURL: h.clientURL,
		DevEnv:    h.devEnv,
	}

	return h.templates.ExecuteTemplate(e.Response(), "product.html", productData)
}
