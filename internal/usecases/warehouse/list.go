package warehouse

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
		Data []WarehouseData `json:"data"`
	}
)

func NewListUsecase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{contextFactory: contextFactory}
}

func (u *listUsecase) Execute(ctx context.Context) (*ListOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	warehouses, err := app.Repositories.Warehouse.List(ctx)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseListError, err)
	}

	data := make([]WarehouseData, 0, len(warehouses))
	for _, w := range warehouses {
		data = append(data, toWarehouseData(w))
	}

	return &ListOutput{Data: data}, nil
}
