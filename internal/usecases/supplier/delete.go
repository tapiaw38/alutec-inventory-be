package supplier

import (
	"context"
	"fmt"

	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
)

type (
	// DeleteUsecase removes a supplier, archives it when only hidden products
	// still point at it, or refuses when products in use would be left
	// pointing at something the user can no longer see.
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

	current, err := app.Repositories.Supplier.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierGetError, err)
	}
	if current == nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierNotFoundError, nil)
	}

	inUse, err := app.Repositories.Product.CountActiveBySupplier(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierDeleteError, err)
	}
	if inUse > 0 {
		return nil, apperrors.NewConflictError(fmt.Sprintf(
			"no se puede eliminar: %s", pluralizeProducts(inUse),
		))
	}

	// Archived products still reference the supplier, so the row has to stay.
	referenced, err := app.Repositories.Product.ExistsBySupplier(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierDeleteError, err)
	}
	if referenced {
		if err := app.Repositories.Supplier.Archive(ctx, id); err != nil {
			return nil, apperrors.NewApplicationError(mappings.SupplierDeleteError, err)
		}
		return &DeleteOutput{Archived: true}, nil
	}

	if err := app.Repositories.Supplier.Delete(ctx, id); err != nil {
		return nil, apperrors.NewApplicationError(mappings.SupplierDeleteError, err)
	}
	return &DeleteOutput{Archived: false}, nil
}

func pluralizeProducts(n int) string {
	if n == 1 {
		return "1 producto lo está usando"
	}
	return fmt.Sprintf("%d productos lo están usando", n)
}
