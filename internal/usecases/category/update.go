package category

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
		Name string `json:"name" binding:"required"`
		Type string `json:"type" binding:"required,oneof=raw_material finished_good tool"`
	}

	UpdateOutput struct {
		Data CategoryData `json:"data"`
	}
)

func NewUpdateUsecase(contextFactory appcontext.Factory) UpdateUsecase {
	return &updateUsecase{contextFactory: contextFactory}
}

func (u *updateUsecase) Execute(ctx context.Context, id string, in UpdateInput) (*UpdateOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	current, err := app.Repositories.Category.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryGetError, err)
	}
	if current == nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryNotFoundError, nil)
	}

	if err := app.Repositories.Category.Update(ctx, id, domain.Category{
		Name: in.Name,
		Type: domain.CategoryType(in.Type),
	}); err != nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryUpdateError, err)
	}

	updated, err := app.Repositories.Category.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryGetError, err)
	}
	if updated == nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryNotFoundError, nil)
	}

	return &UpdateOutput{Data: toCategoryData(*updated)}, nil
}
