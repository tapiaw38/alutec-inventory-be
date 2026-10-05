package stockmovement

import "context"

// ExistsForProduct reports whether the product has any ledger history, which
// decides between deleting it outright and archiving it.
func (r *repository) ExistsForProduct(ctx context.Context, productID string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, `
		SELECT EXISTS (SELECT 1 FROM stock_movements WHERE product_id = $1)
	`, productID).Scan(&exists)
	return exists, err
}
