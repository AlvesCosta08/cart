-- name: CreateProduct :exec
INSERT INTO "produto" (name, price, referencia, categoria_id) VALUES ($1, $2, $3, $4) RETURNING id_produto;

-- name: GetProductByID :one
SELECT * FROM "produto" WHERE id_produto = $1;

-- name: UpdateProduct :exec
UPDATE "produto" SET name = $1, price = $2, referencia = $3, categoria_id = $4 WHERE id_produto = $5;

-- name: DeleteProduct :exec
DELETE FROM "produto" WHERE id_produto = $1;


