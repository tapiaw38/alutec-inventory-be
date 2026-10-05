ALTER TABLE categories ADD COLUMN IF NOT EXISTS archived_at TIMESTAMP;
ALTER TABLE suppliers ADD COLUMN IF NOT EXISTS archived_at TIMESTAMP;
ALTER TABLE warehouses ADD COLUMN IF NOT EXISTS archived_at TIMESTAMP;

-- An archived category keeps its row so archived products still resolve, but
-- it must not keep its name reserved against a new one.
DROP INDEX IF EXISTS idx_categories_name;
CREATE UNIQUE INDEX idx_categories_name ON categories (name) WHERE archived_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_categories_archived_at ON categories (archived_at);
CREATE INDEX IF NOT EXISTS idx_suppliers_archived_at ON suppliers (archived_at);
CREATE INDEX IF NOT EXISTS idx_warehouses_archived_at ON warehouses (archived_at);
