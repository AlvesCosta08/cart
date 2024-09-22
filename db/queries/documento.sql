-- name: CreateDocument :exec
INSERT INTO "documento" (cliente_id, tipo_documento, numero_documento, cpf) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: GetDocumentByClientID :one
SELECT * FROM "documento" WHERE cliente_id = $1;

