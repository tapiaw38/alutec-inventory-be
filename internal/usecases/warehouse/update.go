package warehouse

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
)

type (
	UpdateUsecase interface {
		Execute(ctx context.Context, id string, in UpdateInput) (*UpdateOutput, apperrors.ApplicationError)
	}

	updateUsecase struct {
		contextFactory appcontext.Factory
	}

	UpdateInput struct {
		Name     string `json:"name" binding:"required"`
		Location string `json:"location"`
	}

	UpdateOutput struct {
		Data WarehouseData `json:"data"`
	}
)

func NewUpdateUsecase(contextFactory appcontext.Factory) UpdateUsecase {
	return &updateUsecase{contextFactory: contextFactory}
}

func (u *updateUsecase) Execute(ctx context.Context, id string, in UpdateInput) (*UpdateOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	current, err := app.Repositories.Warehouse.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseGetError, err)
	}
	if current == nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseNotFoundError, nil)
	}

	if err := app.Repositories.Warehouse.Update(ctx, id, domain.Warehouse{
		Name:     in.Name,
		Location: in.Location,
	}); err != nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseUpdateError, err)
	}

	updated, err := app.Repositories.Warehouse.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseGetError, err)
	}
	if updated == nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseNotFoundError, nil)
	}

	return &UpdateOutput{Data: toWarehouseData(*updated)}, nil
}
