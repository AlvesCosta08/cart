-- name: CreateDocument
INSERT INTO "documento" (CPF, CNPJ) VALUES ($1, $2);

-- name: GetDocumentByClientID
SELECT * FROM "documento" WHERE id_cliente = $1;
