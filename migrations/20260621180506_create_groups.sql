-- +goose Up
CREATE TABLE groups (
    id BIGINT PRIMARY KEY,
    parent_id BIGINT UNIQUE NOT NULL,
    name VARCHAR(128) NOT NULL,
    FOREIGN KEY (parent_id) REFERENCES groups(id)
);

-- +goose Down
DROP TABLE groups;
