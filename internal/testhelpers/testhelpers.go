package testhelpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/keto-granola/server/internal/config"
	"github.com/keto-granola/server/internal/server"
	"github.com/labstack/echo/v4"
)

func SetupEchoContext(
	t *testing.T,
	reqBody interface{},
	method string,
	endpoint string,
) (echo.Context, *httptest.ResponseRecorder) {
	t.Helper()

	instance := echo.New()
	server.NewValidator(instance)

	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}

	targetURL := fmt.Sprintf("/%s/%s", config.APIVersion, endpoint)
	var req *http.Request

	switch method {
	case "POST":
		req = httptest.NewRequest(http.MethodPost, targetURL, strings.NewReader(string(jsonBytes)))
	case "GET":
		req = httptest.NewRequest(http.MethodGet, targetURL, http.NoBody)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	ctx := req.Context()
	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()
	return instance.NewContext(req, rec), rec
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
