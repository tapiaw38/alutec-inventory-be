package stockmovement

import (
	"context"
	"time"

	repo "github.com/tapiaw38/alutec-inventory-be/internal/adapters/datasources/repositories/stock_movement"
	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/appcontext"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/errors/mappings"
)

type (
	ListUsecase interface {
		Execute(ctx context.Context, filter ListFilter) (*ListOutput, apperrors.ApplicationError)
	}

	listUsecase struct {
		contextFactory appcontext.Factory
	}

	// ListFilter mirrors the handler's query params, with dates still as
	// strings; it is translated into the repository's ListFilterOptions.
	ListFilter struct {
		ProductID   string
		CategoryID  string
		WarehouseID string
		Type        string
		Search      string
		DateFrom    string
		DateTo      string
	}

	ListOutput struct {
		Data []StockMovementData `json:"data"`
	}
)

func NewListUsecase(contextFactory appcontext.Factory) ListUsecase {
	return &listUsecase{contextFactory: contextFactory}
}

func (u *listUsecase) Execute(ctx context.Context, filter ListFilter) (*ListOutput, apperrors.ApplicationError) {
	app := u.contextFactory()

	options := repo.ListFilterOptions{
		ProductID:   filter.ProductID,
		CategoryID:  filter.CategoryID,
		WarehouseID: filter.WarehouseID,
		Type:        domain.StockMovementType(filter.Type),
		Search:      filter.Search,
	}
	if filter.DateFrom != "" {
		if d, err := time.Parse("2006-01-02", filter.DateFrom); err == nil {
			options.DateFrom = &d
		}
	}
	if filter.DateTo != "" {
		if d, err := time.Parse("2006-01-02", filter.DateTo); err == nil {
			options.DateTo = &d
		}
	}

	movements, err := app.Repositories.StockMovement.List(ctx, options)
	if err != nil {
		return nil, apperrors.NewApplicationError(mappings.StockMovementListError, err)
	}

	data := make([]StockMovementData, 0, len(movements))
	for _, m := range movements {
		data = append(data, toStockMovementData(m))
	}

	return &ListOutput{Data: data}, nil
}
