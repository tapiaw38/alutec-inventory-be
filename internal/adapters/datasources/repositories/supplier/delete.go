package supplier

import "context"

func (r *repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM suppliers WHERE id = $1`, id)
	return err
}
