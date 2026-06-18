package server

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/keto-granola/server/internal/apperr"
)

type Handler[Req any, Res any] func(context.Context, Req) (Res, error)

func Handle[Req any, Res any](handler Handler[Req, Res], statusCode int) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req Req

		if err := ctx.Bind(&req); err != nil {
			return apperr.ToHTTPError(apperr.Validation("request.Validate", "VALIDATION_ERROR", apperr.ErrMsgValidation))
		}

		if err := ctx.Validate(req); err != nil {
			return apperr.ToHTTPError(err)
		}

		res, err := handler(ctx.Request().Context(), req)
		if err != nil {
			return apperr.ToHTTPError(err)
		}

		return ctx.JSON(statusCode, res)
	}
}
