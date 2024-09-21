-- name: CreateTax
INSERT INTO "taxas" (cart_id, value) VALUES ($1, $2) RETURNING id;

-- name: GetTaxByCartID
SELECT * FROM "taxas" WHERE cart_id = $1;

-- name: UpdateTax
UPDATE "taxas" SET value = $1 WHERE id = $2;

-- name: DeleteTax
DELETE FROM "taxas" WHERE id = $1;
