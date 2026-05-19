package apperr

import (
	"fmt"
	"runtime"
	"strings"
)

// Kind categorises errors so callers can branch without string matching.
type Kind uint8

const (
	KindUnknown Kind = iota
	KindNotFound
	KindUnauthorized
	KindValidation
	KindInternal
)

func (k Kind) String() string {
	kinds := [...]string{"unknown", "not_found", "unauthorized", "validation", "internal"}
	if int(k) >= len(kinds) {
		return fmt.Sprintf("kind(%d)", k)
	}
	return kinds[k]
}

type AppError struct {
	Code      string         // machine-readable code, e.g. "PRODUCT_NOT_FOUND"
	Message   string         // human-readable message
	Operation string         // operation, e.g. "ProductService.GetByID"
	Kind      Kind           // error category
	Meta      map[string]any // optional extra context
	Err       error          // wrapped underlying error
	stack     []uintptr
}

func New(operation, code, message string, kind Kind) *AppError {
	stack := make([]uintptr, 32)
	n := runtime.Callers(2, stack)

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

func Unauthorized(operation, message string) *AppError {
	return New(operation, "UNAUTHORIZED", message, KindUnauthorized)
}

func Validation(operation, code, message string) *AppError {
	return New(operation, code, message, KindValidation)
}

func Internal(operation string, err error) *AppError {
	return Wrap(err, operation, "INTERNAL_ERROR", "internal server error", KindInternal)
}
