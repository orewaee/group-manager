-- name: SelectPersonById :one
SELECT * FROM people
WHERE id = $1;

-- name: SelectPersonByGroupId :many
SELECT * FROM people
WHERE group_id = $1;

-- name: DeepSelectPersonByGroupId :many
WITH RECURSIVE subgroups AS (
    SELECT groups.id FROM groups
    WHERE groups.id = $1

    UNION ALL

    SELECT g.id FROM groups g
    INNER JOIN subgroups s ON g.parent_id = s.id
)
SELECT p.* FROM people p
INNER JOIN subgroups s ON p.group_id = s.id;

-- name: InsertPerson :exec
INSERT INTO people (id, firstname, lastname, birthday, group_id)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdatePerson :exec
UPDATE people
SET firstname = $2, lastname = $3, birthday = $4, group_id = $5
WHERE id = $1;

-- name: DeletePerson :exec
DELETE FROM people
WHERE id = $1;
