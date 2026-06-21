-- +goose Up
ALTER TABLE groups ADD CONSTRAINT unique_group_name UNIQUE (name);

-- +goose Down
ALTER TABLE groups DROP CONSTRAINT unique_group_name;
