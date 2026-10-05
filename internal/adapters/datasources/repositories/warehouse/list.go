package warehouse

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) List(ctx context.Context) ([]domain.Warehouse, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, COALESCE(location, ''), created_at
		FROM warehouses
		WHERE archived_at IS NULL
		ORDER BY name ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	warehouses := []domain.Warehouse{}
	for rows.Next() {
		var w domain.Warehouse
		if err := rows.Scan(&w.ID, &w.Name, &w.Location, &w.CreatedAt); err != nil {
			return nil, err
		}
		warehouses = append(warehouses, w)
	}
	return warehouses, rows.Err()
}
