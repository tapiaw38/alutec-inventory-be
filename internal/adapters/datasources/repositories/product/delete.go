package product

import "context"

func (r *repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM products WHERE id = $1`, id)
	return err
}

// Archive hides the product from listings while keeping its row, so the stock
// movements that reference it stay readable.
func (r *repository) Archive(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE products SET archived_at = NOW() WHERE id = $1 AND archived_at IS NULL
	`, id)
	return err
}
