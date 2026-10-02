package supplier

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
		Name  string `json:"name" binding:"required"`
		Phone string `json:"phone"`
		Email string `json:"email"`
	}

	CreateOutput struct {
		Data SupplierData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{contextFactory: contextFactory}
}

func (u *createUsecase) Execute(ctx context.Context, in CreateInput) (*CreateOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	id, err := app.Repositories.Supplier.Create(ctx, domain.Supplier{
		Name:  in.Name,
		Phone: in.Phone,
		Email: in.Email,
	})
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierCreateError, err)
	}

	s, err := app.Repositories.Supplier.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierGetError, err)
	}
	if s == nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierNotFoundError, nil)
	}

	return &CreateOutput{Data: toSupplierData(*s)}, nil
}
