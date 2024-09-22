-- name: CreateOrder :exec
INSERT INTO "pedido" (id_cliente, carrinho_id, status, total) VALUES ($1, $2, $3, $4) RETURNING id_pedido;

-- name: GetOrderByID :one
SELECT * FROM "pedido" WHERE id_pedido = $1;

-- name: UpdateOrder :exec
UPDATE "pedido" SET id_cliente = $1, carrinho_id = $2, status = $3, total = $4 WHERE id_pedido = $5;

-- name: DeleteOrder :exec
DELETE FROM "pedido" WHERE id_pedido = $1;



