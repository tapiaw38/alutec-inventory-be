package product

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Create(ctx context.Context, p domain.Product) (string, error) {
	var id string
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO products (sku, name, category_id, supplier_id, unit, cost_price, sale_price, min_stock, stock_qty, image_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`, p.SKU, p.Name, p.CategoryID, p.SupplierID, p.Unit, p.CostPrice, p.SalePrice, p.MinStock, p.StockQty, p.ImageURL).Scan(&id)
	return id, err
}
