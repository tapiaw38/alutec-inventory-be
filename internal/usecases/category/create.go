package category

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
		Name string `json:"name" binding:"required"`
		Type string `json:"type" binding:"required,oneof=raw_material finished_good tool"`
	}

	CreateOutput struct {
		Data CategoryData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{contextFactory: contextFactory}
}

func (u *createUsecase) Execute(ctx context.Context, in CreateInput) (*CreateOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	id, err := app.Repositories.Category.Create(ctx, domain.Category{
		Name: in.Name,
		Type: domain.CategoryType(in.Type),
	})
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryCreateError, err)
	}

	c, err := app.Repositories.Category.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryGetError, err)
	}
	if c == nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryNotFoundError, nil)
	}

	return &CreateOutput{Data: toCategoryData(*c)}, nil
}
