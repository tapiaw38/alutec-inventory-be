package supplier

import "context"

// Archive hides the record from listings while keeping its row, so anything
// that still references it keeps resolving.
func (r *repository) Archive(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE suppliers SET archived_at = NOW() WHERE id = $1 AND archived_at IS NULL
	`, id)
	return err
}
