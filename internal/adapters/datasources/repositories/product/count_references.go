package product

import "context"

// CountActiveByCategory and CountActiveBySupplier ignore archived products:
// those no longer appear anywhere, so they should not block deleting the
// catalog entry they happen to point at.
func (r *repository) CountActiveByCategory(ctx context.Context, categoryID string) (int, error) {
	return r.countActive(ctx, `category_id = $1`, categoryID)
}

func (r *repository) CountActiveBySupplier(ctx context.Context, supplierID string) (int, error) {
	return r.countActive(ctx, `supplier_id = $1`, supplierID)
}

func (r *repository) countActive(ctx context.Context, condition, id string) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM products WHERE archived_at IS NULL AND `+condition, id,
	).Scan(&count)
	return count, err
}

// ExistsByCategory and ExistsBySupplier include archived products, telling
// "nothing references this" apart from "only hidden things do".
func (r *repository) ExistsByCategory(ctx context.Context, categoryID string) (bool, error) {
	return r.exists(ctx, `category_id = $1`, categoryID)
}

func (r *repository) ExistsBySupplier(ctx context.Context, supplierID string) (bool, error) {
	return r.exists(ctx, `supplier_id = $1`, supplierID)
}

func (r *repository) exists(ctx context.Context, condition, id string) (bool, error) {
	var found bool
	err := r.db.QueryRowContext(ctx,
		`SELECT EXISTS (SELECT 1 FROM products WHERE `+condition+`)`, id,
	).Scan(&found)
	return found, err
}
