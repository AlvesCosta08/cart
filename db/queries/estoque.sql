-- name: CreateStock
INSERT INTO "estoque" (id_produto, quantidade_atual, estoque_minimo) VALUES ($1, $2, $3) RETURNING id_produto;

-- name: GetStockByProductID
SELECT * FROM "estoque" WHERE id_produto = $1;

-- name: UpdateStock
UPDATE "estoque" SET quantidade_atual = $1, estoque_minimo = $2 WHERE id_produto = $3;

-- name: DeleteStock
DELETE FROM "estoque" WHERE id_produto = $1;
