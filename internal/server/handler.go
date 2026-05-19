package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/keto-granola/server/internal/apperr"
)

type Handler[Req any, Res any] func(context.Context, Req) (Res, error)

func Handle[Req any, Res any](handler Handler[Req, Res], statusCode int) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var in Req

		if err := ctx.Bind(&in); err != nil {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				"invalid request",
			)
		}

		out, err := handler(ctx.Request().Context(), in)
		if err != nil {
			return toHTTPError(err)
		}

		return ctx.JSON(statusCode, out)
	}
}

func toHTTPError(err error) *echo.HTTPError {
	var appErr *apperr.AppError
	if !errors.As(err, &appErr) {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error").SetInternal(err)
	}

	switch appErr.Kind {
	case apperr.KindNotFound:
		return echo.NewHTTPError(http.StatusNotFound, appErr.Message)
	case apperr.KindUnauthorized:
		return echo.NewHTTPError(http.StatusUnauthorized, appErr.Message)
	case apperr.KindValidation:
		return echo.NewHTTPError(http.StatusBadRequest, appErr.Message)
	case apperr.KindInternal:
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error").SetInternal(appErr)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error").SetInternal(appErr)
	}
}
