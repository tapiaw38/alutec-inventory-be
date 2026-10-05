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
	// Create takes the warehouse that receives the opening stock; it is
	// ignored when the product starts at zero.
	Create(ctx context.Context, p domain.Product, openingWarehouseID string) (string, error)
	Get(context.Context, string) (*domain.Product, error)
	List(context.Context, ListFilterOptions) ([]domain.Product, error)
	Update(context.Context, string, domain.Product) error
	Delete(context.Context, string) error
	// Archive soft-deletes: the row stays so stock movements keep resolving.
	Archive(ctx context.Context, id string) error
	AdjustStock(ctx context.Context, id string, delta int) error
	// Reference counts decide whether a catalog entry can be removed.
	CountActiveByCategory(ctx context.Context, categoryID string) (int, error)
	CountActiveBySupplier(ctx context.Context, supplierID string) (int, error)
	ExistsByCategory(ctx context.Context, categoryID string) (bool, error)
	ExistsBySupplier(ctx context.Context, supplierID string) (bool, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
