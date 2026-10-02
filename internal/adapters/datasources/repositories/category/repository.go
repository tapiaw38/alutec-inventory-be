package category

import (
	"context"
	"database/sql"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

type Repository interface {
	Create(context.Context, domain.Category) (string, error)
	Get(context.Context, string) (*domain.Category, error)
	List(context.Context) ([]domain.Category, error)
	Update(context.Context, string, domain.Category) error
	Delete(context.Context, string) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
