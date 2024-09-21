-- name: CreateUser
INSERT INTO "user" (nome, email, senha) VALUES ($1, $2, $3) RETURNING id_user;

-- name: GetUserByID
SELECT * FROM "user" WHERE id_user = $1;

-- name: UpdateUser
UPDATE "user" SET nome = $1, email = $2, senha = $3 WHERE id_user = $4;

-- name: DeleteUser
DELETE FROM "user" WHERE id_user = $1;
