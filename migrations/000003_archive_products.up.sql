ALTER TABLE products ADD COLUMN IF NOT EXISTS archived_at TIMESTAMP;

-- An archived product keeps its row for the ledger's sake, but it must not
-- hold its SKU hostage: only active products compete for a SKU.
DROP INDEX IF EXISTS idx_products_sku;
CREATE UNIQUE INDEX idx_products_sku ON products (sku) WHERE archived_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_products_archived_at ON products (archived_at);
