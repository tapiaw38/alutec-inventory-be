package warehouse

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Update(ctx context.Context, id string, w domain.Warehouse) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE warehouses SET name = $1, location = $2 WHERE id = $3
	`, w.Name, w.Location, id)
	return err
}
