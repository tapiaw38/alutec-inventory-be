package stockmovement

import (
	"context"
	"fmt"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) List(ctx context.Context, filter ListFilterOptions) ([]domain.StockMovement, error) {
	query := `
		SELECT sm.id, sm.product_id, sm.warehouse_id, sm.type, sm.quantity, sm.date, COALESCE(sm.note, ''), sm.created_at
		FROM stock_movements sm
		JOIN products p ON p.id = sm.product_id
		WHERE 1 = 1
	`
	args := []any{}

	if filter.ProductID != "" {
		args = append(args, filter.ProductID)
		query += fmt.Sprintf(" AND sm.product_id = $%d", len(args))
	}
	if filter.WarehouseID != "" {
		args = append(args, filter.WarehouseID)
		query += fmt.Sprintf(" AND sm.warehouse_id = $%d", len(args))
	}
	if filter.Type != "" {
		args = append(args, filter.Type)
		query += fmt.Sprintf(" AND sm.type = $%d", len(args))
	}
	if filter.CategoryID != "" {
		args = append(args, filter.CategoryID)
		query += fmt.Sprintf(" AND p.category_id = $%d", len(args))
	}
	if filter.Search != "" {
		args = append(args, "%"+filter.Search+"%")
		query += fmt.Sprintf(" AND (p.name ILIKE $%d OR p.sku ILIKE $%d OR sm.note ILIKE $%d)", len(args), len(args), len(args))
	}
	if filter.DateFrom != nil {
		args = append(args, *filter.DateFrom)
		query += fmt.Sprintf(" AND sm.date >= $%d", len(args))
	}
	if filter.DateTo != nil {
		args = append(args, *filter.DateTo)
		query += fmt.Sprintf(" AND sm.date <= $%d", len(args))
	}
	query += " ORDER BY sm.date DESC, sm.created_at DESC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movements := []domain.StockMovement{}
	for rows.Next() {
		var m domain.StockMovement
		if err := rows.Scan(&m.ID, &m.ProductID, &m.WarehouseID, &m.Type, &m.Quantity, &m.Date, &m.Note, &m.CreatedAt); err != nil {
			return nil, err
		}
		movements = append(movements, m)
	}
	return movements, rows.Err()
}
