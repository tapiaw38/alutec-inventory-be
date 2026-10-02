package product

import (
	"context"
	"fmt"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

// Create inserts the product. When an opening stock is given it is recorded as
// an "in" ledger entry in the same transaction instead of being written
// straight to stock_qty, so every unit in stock is traceable to a movement.
func (r *repository) Create(ctx context.Context, p domain.Product, openingWarehouseID string) (string, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	var id string
	err = tx.QueryRowContext(ctx, `
		INSERT INTO products (sku, name, category_id, supplier_id, unit, cost_price, sale_price, min_stock, stock_qty, image_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 0, $9)
		RETURNING id
	`, p.SKU, p.Name, p.CategoryID, p.SupplierID, p.Unit, p.CostPrice, p.SalePrice, p.MinStock, p.ImageURL).Scan(&id)
	if err != nil {
		return "", err
	}

	if p.StockQty > 0 {
		if _, err := tx.ExecContext(ctx, `
			INSERT INTO stock_movements (product_id, warehouse_id, type, quantity, date, note)
			VALUES ($1, $2, 'in', $3, CURRENT_DATE, 'Stock inicial')
		`, id, openingWarehouseID, p.StockQty); err != nil {
			return "", fmt.Errorf("insert opening stock movement: %w", err)
		}

		if _, err := tx.ExecContext(ctx, `
			UPDATE products SET stock_qty = $1 WHERE id = $2
		`, p.StockQty, id); err != nil {
			return "", fmt.Errorf("set opening stock: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("commit transaction: %w", err)
	}

	return id, nil
}
