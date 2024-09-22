-- name: CreateClient :exec
INSERT INTO "cliente" (nome_fantasia, name, email, telefone, endereco) VALUES ($1, $2, $3, $4, $5) RETURNING id_cliente;

-- name: GetClientByID :one
SELECT * FROM "cliente" WHERE id_cliente = $1;

-- name: UpdateClient :exec
UPDATE "cliente" SET nome_fantasia = $1, name = $2, email = $3, telefone = $4, endereco = $5 WHERE id_cliente = $6;

-- name: DeleteClient :exec
DELETE FROM "cliente" WHERE id_cliente = $1;


