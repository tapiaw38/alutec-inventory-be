package warehouse

import (
	"context"
	"database/sql"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

type Repository interface {
	Create(context.Context, domain.Warehouse) (string, error)
	Get(context.Context, string) (*domain.Warehouse, error)
	List(context.Context) ([]domain.Warehouse, error)
	Update(context.Context, string, domain.Warehouse) error
	Delete(context.Context, string) error
	// Archive soft-deletes: the row stays so existing references resolve.
	Archive(ctx context.Context, id string) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
