package stockmovement

import "context"

// ExistsForWarehouse reports whether the warehouse appears anywhere in the
// ledger, which decides between deleting it and archiving it.
func (r *repository) ExistsForWarehouse(ctx context.Context, warehouseID string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, `
		SELECT EXISTS (SELECT 1 FROM stock_movements WHERE warehouse_id = $1)
	`, warehouseID).Scan(&exists)
	return exists, err
}
