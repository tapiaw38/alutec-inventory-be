package product

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/pgerr"
)

const skuUniqueConstraint = "idx_products_sku"

type (
	CreateUsecase interface {
		Execute(ctx context.Context, in CreateInput) (*CreateOutput, apperrors.ApplicationError)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateInput struct {
		SKU        string  `json:"sku" binding:"required"`
		Name       string  `json:"name" binding:"required"`
		CategoryID string  `json:"category_id" binding:"required"`
		SupplierID string  `json:"supplier_id" binding:"required"`
		Unit       string  `json:"unit"`
		CostPrice  float64 `json:"cost_price"`
		SalePrice  float64 `json:"sale_price"`
		MinStock   int     `json:"min_stock"`
		StockQty   int     `json:"stock_qty"`
		ImageURL   string  `json:"image_url"`
	}

	CreateOutput struct {
		Data ProductData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{contextFactory: contextFactory}
}

func (u *createUsecase) Execute(ctx context.Context, in CreateInput) (*CreateOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	id, err := app.Repositories.Product.Create(ctx, domain.Product{
		SKU:        in.SKU,
		Name:       in.Name,
		CategoryID: in.CategoryID,
		SupplierID: in.SupplierID,
		Unit:       in.Unit,
		CostPrice:  in.CostPrice,
		SalePrice:  in.SalePrice,
		MinStock:   in.MinStock,
		StockQty:   in.StockQty,
		ImageURL:   in.ImageURL,
	})
	if err != nil {
		if pgerr.IsUniqueViolation(err, skuUniqueConstraint) {
			return nil, apperrors.NewConflictError("ya existe un producto con ese SKU")
		}
		if pgerr.IsForeignKeyViolation(err) {
			return nil, apperrors.NewBadRequestError("categoría o proveedor inválido")
		}
		return nil, apperrors.NewApplicationError(mappings.ProductCreateError, err)
	}

	p, err := app.Repositories.Product.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.ProductGetError, err)
	}
	if p == nil {
		return nil, apperrors.NewApplicationError(mappings.ProductNotFoundError, nil)
	}

	return &CreateOutput{Data: toProductData(*p)}, nil
}
