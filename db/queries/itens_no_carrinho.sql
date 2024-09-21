-- name: AddItemToCart
INSERT INTO "itens_no_carrinho" (cart_id, product_id, quantity, price) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: GetItemsByCartID
SELECT * FROM "itens_no_carrinho" WHERE cart_id = $1;

-- name: UpdateCartItem
UPDATE "itens_no_carrinho" SET quantity = $1, price = $2 WHERE id = $3;

-- name: DeleteCartItem
DELETE FROM "itens_no_carrinho" WHERE id = $1;
