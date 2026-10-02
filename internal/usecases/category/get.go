package category

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
		Data CategoryData `json:"data"`
	}
)

func NewGetUsecase(contextFactory appcontext.Factory) GetUsecase {
	return &getUsecase{contextFactory: contextFactory}
}

func (u *getUsecase) Execute(ctx context.Context, id string) (*GetOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	c, err := app.Repositories.Category.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryGetError, err)
	}
	if c == nil {
		return nil, apperrors.NewApplicationError(mappings.CategoryNotFoundError, nil)
	}

	return &GetOutput{Data: toCategoryData(*c)}, nil
}
