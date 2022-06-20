-- name: CreateAnswer :one
INSERT INTO answer (
    text, answer_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetAnswer :one
SELECT * FROM answer
WHERE id = $1 LIMIT 1;

-- name: ListAnswer :many
SELECT * FROM answer
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: GetAnswerForUpdate :one
SELECT * FROM answer
WHERE id = $1 LIMIT 1
FOR UPDATE;

-- name: UpdateAnswer :one
UPDATE answer
SET text = $2
where id = $1
RETURNING *;

-- name: DeleteAnswer :exec
DELETE FROM answer
WHERE id = $1;