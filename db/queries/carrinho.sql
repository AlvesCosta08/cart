-- name: CreateCart
INSERT INTO "carrinho" (user_id) VALUES ($1) RETURNING id;

-- name: GetCartByID
SELECT * FROM "carrinho" WHERE id = $1;

-- name: UpdateCart
UPDATE "carrinho" SET user_id = $1 WHERE id = $2;

-- name: DeleteCart
DELETE FROM "carrinho" WHERE id = $1;
