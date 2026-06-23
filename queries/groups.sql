-- name: SelectGroupById :one
SELECT * FROM groups
WHERE id = $1;

-- name: SelectGroupByName :one
SELECT * FROM groups
WHERE name = $1;
