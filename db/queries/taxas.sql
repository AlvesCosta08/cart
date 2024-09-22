-- name: CreateTax :exec
INSERT INTO "taxas" (cart_id, value) VALUES ($1, $2) RETURNING id;

-- name: GetTaxByCartID :one
SELECT * FROM "taxas" WHERE cart_id = $1;

-- name: UpdateTax :exec
UPDATE "taxas" SET value = $1 WHERE id = $2;

-- name: DeleteTax :exec
DELETE FROM "taxas" WHERE id = $1;

