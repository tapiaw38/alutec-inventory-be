package supplier

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
		Name  string `json:"name" binding:"required"`
		Phone string `json:"phone"`
		Email string `json:"email"`
	}

	UpdateOutput struct {
		Data SupplierData `json:"data"`
	}
)

func NewUpdateUsecase(contextFactory appcontext.Factory) UpdateUsecase {
	return &updateUsecase{contextFactory: contextFactory}
}

func (u *updateUsecase) Execute(ctx context.Context, id string, in UpdateInput) (*UpdateOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	current, err := app.Repositories.Supplier.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierGetError, err)
	}
	if current == nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierNotFoundError, nil)
	}

	if err := app.Repositories.Supplier.Update(ctx, id, domain.Supplier{
		Name:  in.Name,
		Phone: in.Phone,
		Email: in.Email,
	}); err != nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierUpdateError, err)
	}

	updated, err := app.Repositories.Supplier.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierGetError, err)
	}
	if updated == nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierNotFoundError, nil)
	}

	return &UpdateOutput{Data: toSupplierData(*updated)}, nil
}
