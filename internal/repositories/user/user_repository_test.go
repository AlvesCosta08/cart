package user

import (
	sqlc "cart-api/db/sqlc"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)


func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	queries := sqlc.New(db)

	params := sqlc.CreateUserParams{
		Nome:  "João",
		Email: "joao@example.com",
		Senha: "senha123",
	}

	mock.ExpectExec("INSERT INTO \"user\"").
		WithArgs(params.Nome, params.Email, params.Senha).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Chamar o método a ser testado
	err = queries.CreateUser(context.Background(), params)

	// Verificar se não houve erro
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	queries := sqlc.New(db)

	idUser := int32(1)

	mock.ExpectExec("DELETE FROM \"user\"").
		WithArgs(idUser).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Chamar o método a ser testado
	err = queries.DeleteUser(context.Background(), idUser)

	// Verificar se não houve erro
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	queries := sqlc.New(db)

	idUser := int32(1)

	mock.ExpectQuery(`-- name: GetUserByID :one SELECT id_user, nome, email, senha, created_at FROM "user" WHERE id_user = \$1`).
		WithArgs(idUser).
		WillReturnRows(sqlmock.NewRows([]string{"id_user", "nome", "email", "senha", "created_at"}).
			AddRow(idUser, "João", "joao@example.com", "senha123", nil)) // Ajuste aqui se necessário

	// Chamar o método a ser testado
	user, err := queries.GetUserByID(context.Background(), idUser)

	// Verificar se não houve erro e se o usuário retornado está correto
	assert.NoError(t, err)
	assert.Equal(t, sqlc.User{IDUser: idUser, Nome: "João", Email: "joao@example.com", Senha: "senha123"}, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}


func TestGetUserByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	queries := sqlc.New(db)

	idUser := int32(1)

	mock.ExpectQuery(`-- name: GetUserByID :one SELECT id_user, nome, email, senha, created_at FROM "user" WHERE id_user = \$1`).
		WithArgs(idUser).
		WillReturnRows(sqlmock.NewRows([]string{"id_user", "nome", "email", "senha", "created_at"})) // Retorna uma linha vazia

	// Chamar o método a ser testado
	user, err := queries.GetUserByID(context.Background(), idUser)

	// Verificar se houve erro e se o usuário retornado está vazio
	assert.Error(t, err) // Verifica se ocorreu um erro
	assert.Equal(t, sqlc.User{}, user) // Assumindo que você retornaria um User vazio em caso de erro
	assert.NoError(t, mock.ExpectationsWereMet()) // Verifica se todas as expectativas foram atendidas
}


func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	queries := sqlc.New(db)

	params := sqlc.UpdateUserParams{
		IDUser: 1,
		Nome:   "João Atualizado",
		Email:  "joao.atualizado@example.com",
		Senha:  "novaSenha123",
	}

	mock.ExpectExec("UPDATE \"user\"").
		WithArgs(params.Nome, params.Email, params.Senha, params.IDUser).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Chamar o método a ser testado
	err = queries.UpdateUser(context.Background(), params)

	// Verificar se não houve erro
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUser_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	queries := sqlc.New(db)

	params := sqlc.UpdateUserParams{
		IDUser: 1,
		Nome:   "João Atualizado",
		Email:  "joao.atualizado@example.com",
		Senha:  "novaSenha123",
	}

	mock.ExpectExec("UPDATE \"user\"").
		WithArgs(params.Nome, params.Email, params.Senha, params.IDUser).
		WillReturnError(errors.New("error updating user"))

	// Chamar o método a ser testado
	err = queries.UpdateUser(context.Background(), params)

	// Verificar se houve erro
	assert.Error(t, err)
	assert.EqualError(t, err, "error updating user")
	assert.NoError(t, mock.ExpectationsWereMet())
}
