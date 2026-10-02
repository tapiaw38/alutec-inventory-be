package stockmovement

import (
	"context"
	"fmt"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

// Create inserts the ledger entry and applies its stock delta to the product
// in a single transaction, so the ledger and the product balance can never
// drift apart.
func (r *repository) Create(ctx context.Context, m domain.StockMovement) (string, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	var id string
	err = tx.QueryRowContext(ctx, `
		INSERT INTO stock_movements (product_id, warehouse_id, type, quantity, date, note)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`, m.ProductID, m.WarehouseID, m.Type, m.Quantity, m.Date, m.Note).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("insert stock movement: %w", err)
	}

	delta := m.Quantity
	if m.Type == domain.StockMovementTypeOut {
		delta = -m.Quantity
	}

	if _, err := tx.ExecContext(ctx, `
		UPDATE products SET stock_qty = stock_qty + $1 WHERE id = $2
	`, delta, m.ProductID); err != nil {
		return "", fmt.Errorf("adjust product stock: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("commit transaction: %w", err)
	}

	return id, nil
}
