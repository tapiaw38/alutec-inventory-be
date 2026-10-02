package warehouse

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
)

type (
	CreateUsecase interface {
		Execute(ctx context.Context, in CreateInput) (*CreateOutput, apperrors.ApplicationError)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateInput struct {
		Name     string `json:"name" binding:"required"`
		Location string `json:"location"`
	}

	CreateOutput struct {
		Data WarehouseData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{contextFactory: contextFactory}
}

func (u *createUsecase) Execute(ctx context.Context, in CreateInput) (*CreateOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	id, err := app.Repositories.Warehouse.Create(ctx, domain.Warehouse{
		Name:     in.Name,
		Location: in.Location,
	})
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseCreateError, err)
	}

	w, err := app.Repositories.Warehouse.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseGetError, err)
	}
	if w == nil {
		return nil, apperrors.NewApplicationError(mappings.WarehouseNotFoundError, nil)
	}

	return &CreateOutput{Data: toWarehouseData(*w)}, nil
}
