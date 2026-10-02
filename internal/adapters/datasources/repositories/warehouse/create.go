package warehouse

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Create(ctx context.Context, w domain.Warehouse) (string, error) {
	var id string
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO warehouses (name, location)
		VALUES ($1, $2)
		RETURNING id
	`, w.Name, w.Location).Scan(&id)
	return id, err
}
