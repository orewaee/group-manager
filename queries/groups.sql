-- name: SelectGroupById :one
SELECT * FROM groups
WHERE id = $1;

-- name: SelectGroupsByName :many
SELECT * FROM groups
WHERE name = $1;

-- name: SelectAllGroups :many
SELECT * FROM groups;

-- name: SelectChildGroups :many
SELECT * FROM groups
WHERE parent_id = $1;

-- name: CountDirectMembers :one
SELECT COUNT(*) FROM people
WHERE group_id = $1;

-- name: CountTotalMembers :one
WITH RECURSIVE group_tree AS (
    SELECT g.id FROM groups g WHERE g.id = $1
    UNION ALL
    SELECT g.id FROM groups g
    INNER JOIN group_tree gt ON g.parent_id = gt.id
)
SELECT COUNT(*) FROM people p
WHERE p.group_id IN (SELECT gt.id FROM group_tree gt);

-- name: InsertGroup :one
INSERT INTO groups (id, name, parent_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateGroup :exec
UPDATE groups
SET name = $1, parent_id = $2
WHERE id = $3;

-- name: DeleteGroup :exec
DELETE FROM groups
WHERE id = $1;
