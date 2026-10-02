DELETE FROM categories WHERE type = 'tool';

ALTER TABLE categories DROP CONSTRAINT categories_type_check;

ALTER TABLE categories ADD CONSTRAINT categories_type_check
    CHECK (type IN ('raw_material', 'finished_good'));
