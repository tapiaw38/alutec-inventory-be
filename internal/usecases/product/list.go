package product

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/adapters/datasources/repositories/product"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
)

type (
	ListUsecase interface {
		Execute(ctx context.Context, filter ListFilter) (*ListOutput, apperrors.ApplicationError)
	}

	listUsecase struct {
		contextFactory appcontext.Factory
	}

	// ListFilter mirrors the handler's query params; it is translated into
	// the repository's ListFilterOptions, never passed through untouched.
	ListFilter struct {
		CategoryID string
		SupplierID string
		Search     string
		LowStock   bool
	}

	ListOutput struct {
		Data []ProductData `json:"data"`
	}
)

func NewListUsecase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{contextFactory: contextFactory}
}

func (u *listUsecase) Execute(ctx context.Context, filter ListFilter) (*ListOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	products, err := app.Repositories.Product.List(ctx, product.ListFilterOptions{
		CategoryID: filter.CategoryID,
		SupplierID: filter.SupplierID,
		Search:     filter.Search,
		LowStock:   filter.LowStock,
	})
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ProductListError, err)
	}

	data := make([]ProductData, 0, len(products))
	for _, p := range products {
		data = append(data, toProductData(p))
	}

	return &ListOutput{Data: data}, nil
}
