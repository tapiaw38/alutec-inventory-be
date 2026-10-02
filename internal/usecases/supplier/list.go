package supplier

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
)

type (
	ListUsecase interface {
		Execute(ctx context.Context) (*ListOutput, apperrors.ApplicationError)
	}

	listUsecase struct {
		contextFactory appcontext.Factory
	}

	ListOutput struct {
		Data []SupplierData `json:"data"`
	}
)

func NewListUsecase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{contextFactory: contextFactory}
}

func (u *listUsecase) Execute(ctx context.Context) (*ListOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	suppliers, err := app.Repositories.Supplier.List(ctx)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierListError, err)
	}

	data := make([]SupplierData, 0, len(suppliers))
	for _, s := range suppliers {
		data = append(data, toSupplierData(s))
	}

	return &ListOutput{Data: data}, nil
}
