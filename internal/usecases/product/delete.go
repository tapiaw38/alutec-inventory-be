package product

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
)

type (
	// DeleteUsecase removes a product, or archives it when the stock ledger
	// already references it.
	DeleteUsecase interface {
		Execute(ctx context.Context, id string) (*DeleteOutput, apperrors.ApplicationError)
	}

	DeleteOutput struct {
		// Archived is true when history forced a soft delete instead.
		Archived bool `json:"archived"`
	}

	deleteUsecase struct {
		contextFactory appcontext.Factory
	}
)

func NewDeleteUsecase(contextFactory appcontext.Factory) DeleteUsecase {
	return &deleteUsecase{contextFactory: contextFactory}
}

func (u *deleteUsecase) Execute(ctx context.Context, id string) (*DeleteOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	current, err := app.Repositories.Product.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ProductGetError, err)
	}
	if current == nil {
		return nil, apperrors.NewApplicationError(mappings.ProductNotFoundError, nil)
	}

	hasHistory, err := app.Repositories.StockMovement.ExistsForProduct(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ProductDeleteError, err)
	}

	// Deleting a product the ledger points at would take its stock history
	// with it, so those are archived instead.
	if hasHistory {
		if err := app.Repositories.Product.Archive(ctx, id); err != nil {
			return nil, apperrors.NewApplicationError(mappings.ProductDeleteError, err)
		}
		return &DeleteOutput{Archived: true}, nil
	}

	if err := app.Repositories.Product.Delete(ctx, id); err != nil {
		return nil, apperrors.NewApplicationError(mappings.ProductDeleteError, err)
	}

	return &DeleteOutput{Archived: false}, nil
}
