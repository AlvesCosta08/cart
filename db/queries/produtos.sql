-- name: CreateProduct
INSERT INTO "produto" (referencia, descricao, un_medida, preco_unitario, id_categoria) VALUES ($1, $2, $3, $4, $5) RETURNING id_produto;

-- name: GetProductByID
SELECT * FROM "produto" WHERE id_produto = $1;

-- name: UpdateProduct
UPDATE "produto" SET referencia = $1, descricao = $2, un_medida = $3, preco_unitario = $4, id_categoria = $5 WHERE id_produto = $6;

-- name: DeleteProduct
DELETE FROM "produto" WHERE id_produto = $1;
