package supplier

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
)

type (
	GetUsecase interface {
		Execute(ctx context.Context, id string) (*GetOutput, apperrors.ApplicationError)
	}

	getUsecase struct {
		contextFactory appcontext.Factory
	}

	GetOutput struct {
		Data SupplierData `json:"data"`
	}
)

func NewGetUsecase(contextFactory appcontext.Factory) GetUsecase {
	return &getUsecase{contextFactory: contextFactory}
}

func (u *getUsecase) Execute(ctx context.Context, id string) (*GetOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	s, err := app.Repositories.Supplier.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierGetError, err)
	}
	if s == nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierNotFoundError, nil)
	}

	return &GetOutput{Data: toSupplierData(*s)}, nil
}
