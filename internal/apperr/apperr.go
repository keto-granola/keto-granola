package apperr

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/labstack/echo/v4"
)

// Kind categorises errors so callers can branch without string matching.
type Kind uint8

const (
	KindUnknown Kind = iota
	KindNotFound
	KindUnauthorized
	KindValidation
	KindInternal

	callerDepth    = 2
	maxStackFrames = 32

	ErrMsgValidation = "invalid request"
	ErrMsgInternal   = "internal server error"
)

func (k Kind) String() string {
	kinds := [...]string{"unknown", "not_found", "unauthorised", "validation", "internal"}
	if int(k) >= len(kinds) {
		return fmt.Sprintf("kind(%d)", k)
	}
	return kinds[k]
}

type AppError struct {
	Code      string // machine-readable code, e.g. "PRODUCT_NOT_FOUND"
	Message   string // human-readable message
	Operation string // operation, e.g. "ProductService.GetByID"
	Kind      Kind   // error category
	Err       error  // wrapped underlying error
	stack     []uintptr
}

func New(operation, code, message string, kind Kind) *AppError {
	stack := make([]uintptr, maxStackFrames)
	n := runtime.Callers(callerDepth, stack)

	return &AppError{
		Operation: operation,
		Code:      code,
		Message:   message,
		Kind:      kind,
		stack:     stack[:n],
	}
}

func Wrap(err error, operation, code, message string, kind Kind) *AppError {
	e := New(operation, code, message, kind)
	e.Err = err

	return e
}

func (e *AppError) Error() string {
	var sb strings.Builder

	if e.Operation != "" {
		fmt.Fprintf(&sb, "[%s] ", e.Operation)
	}

	fmt.Fprintf(&sb, "%s: %s", e.Code, e.Message)
	if e.Err != nil {
		fmt.Fprintf(&sb, ": %v", e.Err)
	}

	return sb.String()
}

func (e *AppError) StackTrace() string {
	frames := runtime.CallersFrames(e.stack)
	var sb strings.Builder

	for {
		f, more := frames.Next()
		fmt.Fprintf(&sb, "%s\n\t%s:%d\n", f.Function, f.File, f.Line)
		if !more {
			break
		}
	}

	return sb.String()
}

func NotFound(operation, code, message string) *AppError {
	return New(operation, code, message, KindNotFound)
}

func Unauthorised(operation, message string) *AppError {
	return New(operation, "UNAUTHORISED", message, KindUnauthorized)
}

func Validation(operation, code, message string) *AppError {
	return New(operation, code, message, KindValidation)
}

func Internal(operation string, err error) *AppError {
	return Wrap(err, operation, "INTERNAL_ERROR", ErrMsgInternal, KindInternal)
}

func ToHTTPError(err error) *echo.HTTPError {
	var appErr *AppError
	if !errors.As(err, &appErr) {
		return echo.NewHTTPError(http.StatusInternalServerError, ErrMsgInternal).SetInternal(err)
	}

	switch appErr.Kind {
	case KindNotFound:
		return echo.NewHTTPError(http.StatusNotFound, appErr.Message)
	case KindUnauthorized:
		return echo.NewHTTPError(http.StatusUnauthorized, appErr.Message)
	case KindValidation:
		return echo.NewHTTPError(http.StatusBadRequest, appErr.Message)
	case KindInternal:
		return echo.NewHTTPError(http.StatusInternalServerError, ErrMsgInternal).SetInternal(appErr)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, ErrMsgInternal).SetInternal(appErr)
	}
}
