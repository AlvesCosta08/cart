-- name: CreateClient
INSERT INTO "cliente" (nome_fantasia, razao_social, tipo_cliente, cpf, cnpj, inscricao_estadual, email, endereco, cep, cidade, uf, fone) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id_cliente;

-- name: GetClientByID
SELECT * FROM "cliente" WHERE id_cliente = $1;

-- name: UpdateClient
UPDATE "cliente" SET nome_fantasia = $1, razao_social = $2, tipo_cliente = $3, cpf = $4, cnpj = $5, inscricao_estadual = $6, email = $7, endereco = $8, cep = $9, cidade = $10, uf = $11, fone = $12 WHERE id_cliente = $13;

-- name: DeleteClient
DELETE FROM "cliente" WHERE id_cliente = $1;
