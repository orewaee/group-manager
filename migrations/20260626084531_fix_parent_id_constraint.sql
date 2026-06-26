-- +goose Up
ALTER TABLE groups ALTER COLUMN parent_id DROP NOT NULL;
ALTER TABLE groups DROP CONSTRAINT IF EXISTS groups_parent_id_key;

-- +goose Down
ALTER TABLE groups ALTER COLUMN parent_id SET NOT NULL;
