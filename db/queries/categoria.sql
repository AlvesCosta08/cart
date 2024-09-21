-- name: CreateCategory
INSERT INTO "categoria" (nome, sub_categoria) VALUES ($1, $2) RETURNING id_categoria;

-- name: GetCategoryByID
SELECT * FROM "categoria" WHERE id_categoria = $1;

-- name: UpdateCategory
UPDATE "categoria" SET nome = $1, sub_categoria = $2 WHERE id_categoria = $3;

-- name: DeleteCategory
DELETE FROM "categoria" WHERE id_categoria = $1;
