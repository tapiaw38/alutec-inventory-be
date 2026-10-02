package product

import (
	"context"
	"database/sql"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

type ListFilterOptions struct {
	CategoryID string
	SupplierID string
	Search     string
	LowStock   bool
}

type Repository interface {
	Create(context.Context, domain.Product) (string, error)
	Get(context.Context, string) (*domain.Product, error)
	List(context.Context, ListFilterOptions) ([]domain.Product, error)
	Update(context.Context, string, domain.Product) error
	Delete(context.Context, string) error
	AdjustStock(ctx context.Context, id string, delta int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
