DROP INDEX IF EXISTS idx_warehouses_archived_at;
DROP INDEX IF EXISTS idx_suppliers_archived_at;
DROP INDEX IF EXISTS idx_categories_archived_at;

-- Archived names would collide once the index covers every category again.
DELETE FROM categories WHERE archived_at IS NOT NULL;

DROP INDEX IF EXISTS idx_categories_name;
CREATE UNIQUE INDEX idx_categories_name ON categories (name);

ALTER TABLE warehouses DROP COLUMN IF EXISTS archived_at;
ALTER TABLE suppliers DROP COLUMN IF EXISTS archived_at;
ALTER TABLE categories DROP COLUMN IF EXISTS archived_at;
