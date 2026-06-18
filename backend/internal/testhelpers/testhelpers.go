package testhelpers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/keto-granola/server/internal/config"
	"github.com/keto-granola/server/internal/server"
)

func SetupEchoContext(
	t *testing.T,
	reqBody interface{},
	method string,
	endpoint string,
) (echo.Context, *httptest.ResponseRecorder) {
	t.Helper()
	ctx := context.Background()

	e := echo.New()
	server.NewValidator(e)

	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}

	targetURL := fmt.Sprintf("%s/%s", config.APIBasePath, endpoint)
	var req *http.Request

	switch method {
	case http.MethodPost:
		req = httptest.NewRequestWithContext(ctx, method, targetURL, bytes.NewReader(jsonBytes))
	case http.MethodPatch:
		req = httptest.NewRequestWithContext(ctx, method, targetURL, bytes.NewReader(jsonBytes))
	case http.MethodGet:
		req = httptest.NewRequestWithContext(ctx, method, targetURL, http.NoBody)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func AssertHTTPError(t *testing.T, err error, expectedCode int, expectedMsg string) {
	t.Helper()

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	httpErr, ok := err.(*echo.HTTPError)
	if !ok {
		t.Fatalf("expected *echo.HTTPError, got %T", err)
	}

	if httpErr.Code != expectedCode {
		t.Errorf("expected status %d, got %d", expectedCode, httpErr.Code)
	}

	if httpErr.Message != expectedMsg {
		t.Errorf("expected message %q, got %v", expectedMsg, httpErr.Message)
	}
}

func AssertRepoCalls(t *testing.T, got, expected int, handlerName string) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %d calls to %s, got %d", expected, handlerName, got)
	}
}
