-- name: CreateCart :exec
INSERT INTO "carrinho" (user_id) VALUES ($1) RETURNING id;

-- name: GetCartByID :one
SELECT * FROM "carrinho" WHERE id = $1;

-- name: UpdateCart :exec
UPDATE "carrinho" SET user_id = $1 WHERE id = $2;

-- name: DeleteCart :exec
DELETE FROM "carrinho" WHERE id = $1;

