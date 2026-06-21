-- +goose Up
CREATE TABLE people (
    id BIGINT PRIMARY KEY,
    firstname VARCHAR(32) NOT NULL,
    lastname VARCHAR(32) NOT NULL,
    birthday DATE NOT NULL DEFAULT CURRENT_DATE,
    group_id BIGINT NOT NULL,

    FOREIGN KEY (group_id) REFERENCES groups(id)
);

-- +goose Down
DROP TABLE people;
