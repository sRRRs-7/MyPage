-- name: CreateBlog :one
INSERT INTO blog (
    title, text, image
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetBlog :one
SELECT * FROM blog
WHERE id = $1 LIMIT 1;

-- name: ListBlog :many
SELECT * FROM blog
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: GetBlogForUpdate :one
SELECT * FROM blog
WHERE id = $1 LIMIT 1
FOR UPDATE;

-- name: UpdateBlog :one
UPDATE blog
SET title = $2, text = $3, image = $4
where id = $1
RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blog
WHERE id = $1;