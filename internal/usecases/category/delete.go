package category

import (
	"context"
	"fmt"

	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
)

type (
	// DeleteUsecase removes a category, archives it when only hidden products
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

	current, err := app.Repositories.Category.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryGetError, err)
	}
	if current == nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryNotFoundError, nil)
	}

	inUse, err := app.Repositories.Product.CountActiveByCategory(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryDeleteError, err)
	}
	if inUse > 0 {
		return nil, apperrors.NewConflictError(fmt.Sprintf(
			"no se puede eliminar: %s", pluralizeProducts(inUse),
		))
	}

	// Archived products still reference the category, so the row has to stay.
	referenced, err := app.Repositories.Product.ExistsByCategory(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryDeleteError, err)
	}
	if referenced {
		if err := app.Repositories.Category.Archive(ctx, id); err != nil {
			return nil, apperrors.NewApplicationError(mappings.CategoryDeleteError, err)
		}
		return &DeleteOutput{Archived: true}, nil
	}

	if err := app.Repositories.Category.Delete(ctx, id); err != nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryDeleteError, err)
	}
	return &DeleteOutput{Archived: false}, nil
}

func pluralizeProducts(n int) string {
	if n == 1 {
		return "1 producto la está usando"
	}
	return fmt.Sprintf("%d productos la están usando", n)
}
