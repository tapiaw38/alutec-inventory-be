package warehouse

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
)

type (
	// DeleteUsecase removes a warehouse, or archives it once the stock ledger
	// references it. Unlike products, a warehouse is never blocked: the ledger
	// is history, not something the user can clear out first.
	DeleteUsecase interface {
		Execute(ctx context.Context, id string) (*DeleteOutput, apperrors.ApplicationError)
	}

	deleteUsecase struct {
		contextFactory appcontext.Factory
	}

	DeleteOutput struct {
		Archived bool `json:"archived"`
	}
)

func NewDeleteUsecase(contextFactory appcontext.Factory) DeleteUsecase {
	return &deleteUsecase{contextFactory: contextFactory}
}

func (u *deleteUsecase) Execute(ctx context.Context, id string) (*DeleteOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	current, err := app.Repositories.Warehouse.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseGetError, err)
	}
	if current == nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseNotFoundError, nil)
	}

	referenced, err := app.Repositories.StockMovement.ExistsForWarehouse(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseDeleteError, err)
	}
	if referenced {
		if err := app.Repositories.Warehouse.Archive(ctx, id); err != nil {
			return nil, apperrors.NewApplicationError(mappings.WarehouseDeleteError, err)
		}
		return &DeleteOutput{Archived: true}, nil
	}

	if err := app.Repositories.Warehouse.Delete(ctx, id); err != nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseDeleteError, err)
	}
	return &DeleteOutput{Archived: false}, nil
}
