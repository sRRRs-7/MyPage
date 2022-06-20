-- name: CreateQuestion :one
INSERT INTO question (
    text
) VALUES (
    $1
) RETURNING *;

-- name: GetQuestion :one
SELECT * FROM question
WHERE id = $1 LIMIT 1;

-- name: ListQuestion :many
SELECT * FROM question
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: GetQuestionForUpdate :one
SELECT * FROM question
WHERE id = $1 LIMIT 1
FOR UPDATE;

-- name: UpdateQuestion :one
UPDATE question
SET text = $2
where id = $1
RETURNING *;

-- name: DeleteQuestion :exec
DELETE FROM question
WHERE id = $1;