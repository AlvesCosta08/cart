-- name: AddItemToCart :exec
INSERT INTO "itens_no_carrinho" (cart_id, produto_id, quantidade, preco_unitario) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: GetItemsByCartID :many
SELECT * FROM "itens_no_carrinho" WHERE cart_id = $1;

-- name: UpdateCartItem :exec
UPDATE "itens_no_carrinho" SET quantidade = $1, preco_unitario = $2 WHERE id = $3;

-- name: DeleteCartItem :exec
DELETE FROM "itens_no_carrinho" WHERE id = $1;


