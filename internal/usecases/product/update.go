package product

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/pgerr"
)

type (
	UpdateUsecase interface {
		Execute(ctx context.Context, id string, in UpdateInput) (*UpdateOutput, apperrors.ApplicationError)
	}

	updateUsecase struct {
		contextFactory appcontext.Factory
	}

	UpdateInput struct {
		SKU        string  `json:"sku" binding:"required"`
		Name       string  `json:"name" binding:"required"`
		CategoryID string  `json:"category_id" binding:"required"`
		SupplierID string  `json:"supplier_id" binding:"required"`
		Unit       string  `json:"unit"`
		CostPrice  float64 `json:"cost_price"`
		SalePrice  float64 `json:"sale_price"`
		MinStock   int     `json:"min_stock"`
		ImageURL   string  `json:"image_url"`
	}

	UpdateOutput struct {
		Data ProductData `json:"data"`
	}
)

func NewUpdateUsecase(contextFactory appcontext.Factory) UpdateUsecase {
	return &updateUsecase{contextFactory: contextFactory}
}

func (u *updateUsecase) Execute(ctx context.Context, id string, in UpdateInput) (*UpdateOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	current, err := app.Repositories.Product.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ProductGetError, err)
	}
	if current == nil {
		return nil, apperrors.NewApplicationError(mappings.ProductNotFoundError, nil)
	}

	if err := app.Repositories.Product.Update(ctx, id, domain.Product{
		SKU:        in.SKU,
		Name:       in.Name,
		CategoryID: in.CategoryID,
		SupplierID: in.SupplierID,
		Unit:       in.Unit,
		CostPrice:  in.CostPrice,
		SalePrice:  in.SalePrice,
		MinStock:   in.MinStock,
		ImageURL:   in.ImageURL,
	}); err != nil {
		if pgerr.IsUniqueViolation(err, skuUniqueConstraint) {
			return nil, apperrors.NewConflictError("ya existe un producto con ese SKU")
		}
		if pgerr.IsForeignKeyViolation(err) {
			return nil, apperrors.NewBadRequestError("categoría o proveedor inválido")
		}
		return nil, apperrors.NewApplicationError(mappings.ProductUpdateError, err)
	}

	updated, err := app.Repositories.Product.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ProductGetError, err)
	}
	if updated == nil {
		return nil, apperrors.NewApplicationError(mappings.ProductNotFoundError, nil)
	}

	return &UpdateOutput{Data: toProductData(*updated)}, nil
}
