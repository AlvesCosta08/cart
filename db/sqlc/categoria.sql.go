// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: categoria.sql

package db

import (
	"context"
	"database/sql"
)

const createCategory = `-- name: CreateCategory :exec
INSERT INTO "categoria" (nome, description) VALUES ($1, $2) RETURNING id_categoria
`

type CreateCategoryParams struct {
	Nome        string
	Description sql.NullString
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory, arg.Nome, arg.Description)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM "categoria" WHERE id_categoria = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, idCategoria int32) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, idCategoria)
	return err
}

const getCategoryByID = `-- name: GetCategoryByID :one
SELECT id_categoria, nome, description, created_at FROM "categoria" WHERE id_categoria = $1
`

func (q *Queries) GetCategoryByID(ctx context.Context, idCategoria int32) (Categorium, error) {
	row := q.db.QueryRowContext(ctx, getCategoryByID, idCategoria)
	var i Categorium
	err := row.Scan(
		&i.IDCategoria,
		&i.Nome,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE "categoria" SET nome = $1, description = $2 WHERE id_categoria = $3
`

type UpdateCategoryParams struct {
	Nome        string
	Description string	
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory, arg.Nome, arg.Description)
	return err
}
