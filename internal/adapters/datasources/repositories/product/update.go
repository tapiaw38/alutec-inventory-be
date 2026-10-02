package product

import (
	"context"

	"github.com/tapiaw38/alutec-inventory-be/internal/domain"
)

func (r *repository) Update(ctx context.Context, id string, p domain.Product) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE products
		SET sku = $1, name = $2, category_id = $3, supplier_id = $4, unit = $5,
		    cost_price = $6, sale_price = $7, min_stock = $8, image_url = $9
		WHERE id = $10
	`, p.SKU, p.Name, p.CategoryID, p.SupplierID, p.Unit, p.CostPrice, p.SalePrice, p.MinStock, p.ImageURL, id)
	return err
}
