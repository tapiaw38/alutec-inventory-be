package stockmovement

import (
	"context"
	"database/sql"
	"time"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

// ListFilterOptions describes query mechanics only; it is not a domain entity.
type ListFilterOptions struct {
	ProductID   string
	CategoryID  string
	WarehouseID string
	Type        domain.StockMovementType
	Search      string
	DateFrom    *time.Time
	DateTo      *time.Time
}

// Repository is an append-only ledger: there is deliberately no Update or
// Delete. To correct a mistake, create a new movement that reverses it.
type Repository interface {
	Create(context.Context, domain.StockMovement) (string, error)
	Get(context.Context, string) (*domain.StockMovement, error)
	List(context.Context, ListFilterOptions) ([]domain.StockMovement, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
