DROP INDEX IF EXISTS idx_products_archived_at;

-- Archived rows would collide once the SKU index covers every product again.
DELETE FROM products WHERE archived_at IS NOT NULL;

DROP INDEX IF EXISTS idx_products_sku;
CREATE UNIQUE INDEX idx_products_sku ON products (sku);

ALTER TABLE products DROP COLUMN IF EXISTS archived_at;
