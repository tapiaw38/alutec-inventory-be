package product

import "context"

func (r *repository) AdjustStock(ctx context.Context, id string, delta int) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE products SET stock_qty = stock_qty + $1 WHERE id = $2
	`, delta, id)
	return err
}
