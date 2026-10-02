package supplier

import (
	"context"
	"database/sql"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

type Repository interface {
	Create(context.Context, domain.Supplier) (string, error)
	Get(context.Context, string) (*domain.Supplier, error)
	List(context.Context) ([]domain.Supplier, error)
	Update(context.Context, string, domain.Supplier) error
	Delete(context.Context, string) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
