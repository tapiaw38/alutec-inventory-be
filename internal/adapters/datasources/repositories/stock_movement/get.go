package stockmovement

import (
	"context"
	"database/sql"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Get(ctx context.Context, id string) (*domain.StockMovement, error) {
	var m domain.StockMovement
	err := r.db.QueryRowContext(ctx, `
		SELECT id, product_id, warehouse_id, type, quantity, date, COALESCE(note, ''), created_at
		FROM stock_movements
		WHERE id = $1
	`, id).Scan(&m.ID, &m.ProductID, &m.WarehouseID, &m.Type, &m.Quantity, &m.Date, &m.Note, &m.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &m, nil
}
