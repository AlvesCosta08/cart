-- name: CreateCategory :exec
INSERT INTO "categoria" (nome, description) VALUES ($1, $2) RETURNING id_categoria;

-- name: GetCategoryByID :one
SELECT * FROM "categoria" WHERE id_categoria = $1;

-- name: UpdateCategory :exec
UPDATE "categoria" SET nome = $1, description = $2 WHERE id_categoria = $3;

-- name: DeleteCategory :exec
DELETE FROM "categoria" WHERE id_categoria = $1;


