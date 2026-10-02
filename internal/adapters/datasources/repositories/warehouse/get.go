package warehouse

import (
	"context"
	"database/sql"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Get(ctx context.Context, id string) (*domain.Warehouse, error) {
	var w domain.Warehouse
	err := r.db.QueryRowContext(ctx, `
		SELECT id, name, COALESCE(location, ''), created_at
		FROM warehouses
		WHERE id = $1
	`, id).Scan(&w.ID, &w.Name, &w.Location, &w.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &w, nil
}
