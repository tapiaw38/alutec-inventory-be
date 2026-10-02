package warehouse

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
)

type (
	DeleteUsecase interface {
		Execute(ctx context.Context, id string) apperrors.ApplicationError
	}

	deleteUsecase struct {
		contextFactory appcontext.Factory
	}
)

func NewDeleteUsecase(contextFactory appcontext.Factory) DeleteUsecase {
	return &deleteUsecase{contextFactory: contextFactory}
}

func (u *deleteUsecase) Execute(ctx context.Context, id string) apperrors.ApplicationError {
	app := u.contextFactory()

	current, err := app.Repositories.Warehouse.Get(ctx, id)
	if err != nil {
		return apperrors.NewApplicationError(mappings.WarehouseGetError, err)
	}
	if current == nil {
		return apperrors.NewApplicationError(mappings.WarehouseNotFoundError, nil)
	}

	if err := app.Repositories.Warehouse.Delete(ctx, id); err != nil {
		return apperrors.NewApplicationError(mappings.WarehouseDeleteError, err)
	}

	return nil
}
