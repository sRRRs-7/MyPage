// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: blog.sql

package db

import (
	"context"
)

const createBlog = `-- name: CreateBlog :one
INSERT INTO blog (
    title, text, image
) VALUES (
    $1, $2, $3
) RETURNING id, title, text, image, created_at, updated_at
`

type CreateBlogParams struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Image []byte `json:"image"`
}

func (q *Queries) CreateBlog(ctx context.Context, arg CreateBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, createBlog, arg.Title, arg.Text, arg.Image)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Text,
		&i.Image,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteBlog = `-- name: DeleteBlog :exec
DELETE FROM blog
WHERE id = $1
`

func (q *Queries) DeleteBlog(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBlog, id)
	return err
}

const getBlog = `-- name: GetBlog :one
SELECT id, title, text, image, created_at, updated_at FROM blog
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBlog(ctx context.Context, id int64) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getBlog, id)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Text,
		&i.Image,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBlogForUpdate = `-- name: GetBlogForUpdate :one
SELECT id, title, text, image, created_at, updated_at FROM blog
WHERE id = $1 LIMIT 1
FOR UPDATE
`

func (q *Queries) GetBlogForUpdate(ctx context.Context, id int64) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getBlogForUpdate, id)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Text,
		&i.Image,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listBlog = `-- name: ListBlog :many
SELECT id, title, text, image, created_at, updated_at FROM blog
ORDER BY id DESC
LIMIT $1
OFFSET $2
`

type ListBlogParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBlog(ctx context.Context, arg ListBlogParams) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, listBlog, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Blog{}
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Text,
			&i.Image,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBlog = `-- name: UpdateBlog :one
UPDATE blog
SET title = $2, text = $3, image = $4
where id = $1
RETURNING id, title, text, image, created_at, updated_at
`

type UpdateBlogParams struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Image []byte `json:"image"`
}

func (q *Queries) UpdateBlog(ctx context.Context, arg UpdateBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, updateBlog,
		arg.ID,
		arg.Title,
		arg.Text,
		arg.Image,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Text,
		&i.Image,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
