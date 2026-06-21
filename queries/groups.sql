-- name: FindById :one
SELECT * FROM groups
WHERE id = $1;

-- name: FindByName :one
SELECT * FROM groups
WHERE name = $1;
