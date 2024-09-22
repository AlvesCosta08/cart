-- name: CreateStock :exec
INSERT INTO "estoque" (id_produto, quantidade_atual) VALUES ($1, $2) RETURNING id;

-- name: GetStockByProductID :one
SELECT * FROM "estoque" WHERE id_produto = $1;

-- name: UpdateStock :exec
UPDATE "estoque" SET quantidade_atual = $1 WHERE id_produto = $2;

-- name: DeleteStock :exec
DELETE FROM "estoque" WHERE id_produto = $1;


