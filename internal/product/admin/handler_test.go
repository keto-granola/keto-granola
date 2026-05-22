package admin_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/keto-granola/server/internal/product"
	"github.com/keto-granola/server/internal/product/admin"
	"github.com/keto-granola/server/internal/product/mocks"
	"github.com/keto-granola/server/internal/server"
	"github.com/keto-granola/server/internal/testhelpers"
)

func TestCreateProduct_HappyPath(t *testing.T) {
	t.Run("returns new product successfully", func(t *testing.T) {
		mockRepo := mocks.RepositoryMock{
			InsertProductFunc: func(ctx context.Context, params *product.CreateProductParams) (*product.Product, error) {
				if diff := cmp.Diff(expCreateProdParams, params); diff != "" {
					t.Errorf("repo received wrong params (-want +got):\n%s", diff)
				}

				return insertedProd, nil
			},
		}

		ctx, rec := testhelpers.SetupEchoContext(t, createProdReqBody, http.MethodPost, "/admin/products")

		h := admin.NewHandler(admin.NewService(&mockRepo))

		handlerFunc := server.Handle(h.CreateProduct, http.StatusCreated)
		err := handlerFunc(ctx)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if rec.Code != http.StatusCreated {
			t.Errorf("expected status %d, got %d", http.StatusCreated, rec.Code)
		}

		var actual admin.CreateProductResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
			t.Fatalf("unmarshal response: %v", err)
		}

		if diff := cmp.Diff(expCreateProdRes, actual); diff != "" {
			t.Errorf("product mismatch (-want +got):\n%s", diff)
		}

		testhelpers.AssertRepoCalls(t, len(mockRepo.InsertProductCalls()), 1, "admin.InsertProduct")
	})
}

func TestCreateProduct_UnhappyPath(t *testing.T) {
	for _, tt := range createProdUnhappyPathTestCases {
		t.Run(tt.name, func(t *testing.T) {
			h := tt.arrange()

			ctx, _ := testhelpers.SetupEchoContext(t, tt.reqBody, http.MethodPost, "/admin/products")

			handlerFunc := server.Handle(h.CreateProduct, http.StatusCreated)
			err := handlerFunc(ctx)

			testhelpers.AssertHTTPError(t, err, tt.wantHTTPStatus, tt.expectedErrMsg)
		})
	}
}
