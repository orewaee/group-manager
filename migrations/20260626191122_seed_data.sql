-- +goose Up
INSERT INTO groups (id, parent_id, name) VALUES
    (1, NULL, 'Engineering'),
    (2, 1,    'Backend'),
    (3, 1,    'Frontend'),
    (4, 1,    'QA'),
    (5, NULL, 'Design');

INSERT INTO people (id, firstname, lastname, birthday, group_id, created_at, updated_at) VALUES
    (1, 'Alice',   'Smith',    '1990-03-15', 1, NOW(), NOW()),
    (2, 'Bob',     'Johnson',  '1985-07-22', 2, NOW(), NOW()),
    (3, 'Charlie', 'Brown',    '1992-11-08', 2, NOW(), NOW()),
    (4, 'Diana',   'Prince',   '1988-05-30', 3, NOW(), NOW()),
    (5, 'Eve',     'Adams',    '1995-09-12', 3, NOW(), NOW()),
    (6, 'Frank',   'Castle',   '1982-01-25', 4, NOW(), NOW()),
    (7, 'Grace',   'Lee',      '1993-12-01', 5, NOW(), NOW());

-- +goose Down
DELETE FROM people WHERE id BETWEEN 1 AND 7;
DELETE FROM groups WHERE id BETWEEN 1 AND 5;
