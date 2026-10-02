package stockmovement

import (
	"context"
	"time"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/pgerr"
)

type (
	// CreateUsecase is the only way to affect stock levels: there is no
	// separate "update product stock" operation. Selling, restocking, or
	// correcting a mistake are all a new ledger entry.
	CreateUsecase interface {
		Execute(ctx context.Context, in CreateInput) (*CreateOutput, apperrors.ApplicationError)
	}

	createUsecase struct {
		contextFactory appcontext.Factory
	}

	CreateInput struct {
		ProductID   string `json:"product_id" binding:"required"`
		WarehouseID string `json:"warehouse_id" binding:"required"`
		Type        string `json:"type" binding:"required,oneof=in out adjustment"`
		Quantity    int    `json:"quantity" binding:"required,gt=0"`
		Date        string `json:"date" binding:"required"`
		Note        string `json:"note"`
	}

	CreateOutput struct {
		Data StockMovementData `json:"data"`
	}
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase {
	return &createUsecase{contextFactory: contextFactory}
}

func (u *createUsecase) Execute(ctx context.Context, in CreateInput) (*CreateOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	date, parseErr := time.Parse("2006-01-02", in.Date)
	if parseErr != nil {
		return nil, apperrors.NewBadRequestError("date must be in YYYY-MM-DD format")
	}

	id, err := app.Repositories.StockMovement.Create(ctx, domain.StockMovement{
		ProductID:   in.ProductID,
		WarehouseID: in.WarehouseID,
		Type:        domain.StockMovementType(in.Type),
		Quantity:    in.Quantity,
		Date:        date,
		Note:        in.Note,
	})
	if err != nil {
		if pgerr.IsForeignKeyViolation(err) {
			return nil, apperrors.NewBadRequestError("producto o depósito inválido")
		}
		return nil, apperrors.NewApplicationError(mappings.StockMovementCreateError, err)
	}

	m, err := app.Repositories.StockMovement.Get(ctx, id)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.StockMovementGetError, err)
	}
	if m == nil {
		return nil, apperrors.NewApplicationError(mappings.StockMovementNotFoundError, nil)
	}

	return &CreateOutput{Data: toStockMovementData(*m)}, nil
}
