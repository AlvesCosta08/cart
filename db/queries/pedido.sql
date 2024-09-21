-- name: CreateOrder
INSERT INTO "pedido" (id_cliente, cart_id, forma_pagamento, total) VALUES ($1, $2, $3, $4) RETURNING id_pedido;

-- name: GetOrderByID
SELECT * FROM "pedido" WHERE id_pedido = $1;

-- name: UpdateOrder
UPDATE "pedido" SET id_cliente = $1, cart_id = $2, forma_pagamento = $3, total = $4 WHERE id_pedido = $5;

-- name: DeleteOrder
DELETE FROM "pedido" WHERE id_pedido = $1;
