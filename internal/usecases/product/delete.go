package product

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

	current, err := app.Repositories.Product.Get(ctx, id)
	if err != nil {
		return apperrors.NewApplicationError(mappings.ProductGetError, err)
	}
	if current == nil {
		return apperrors.NewApplicationError(mappings.ProductNotFoundError, nil)
	}

	if err := app.Repositories.Product.Delete(ctx, id); err != nil {
		return apperrors.NewApplicationError(mappings.ProductDeleteError, err)
	}

	return nil
}
