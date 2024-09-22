-- name: CreateUser :exec
INSERT INTO "user" (nome, email, senha) VALUES ($1, $2, $3) RETURNING id_user;

-- name: GetUserByID :one
SELECT * FROM "user" WHERE id_user = $1;

-- name: UpdateUser :exec
UPDATE "user" SET nome = $1, email = $2, senha = $3 WHERE id_user = $4;

-- name: DeleteUser :exec
DELETE FROM "user" WHERE id_user = $1;

