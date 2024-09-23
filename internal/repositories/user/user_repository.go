package user

import (
	"context"

	db "cart-api/db/sqlc"
)

// UserRepository define os métodos para interagir com a tabela "user".
type UserRepository interface {
	CreateUser(ctx context.Context, params db.CreateUserParams) error
	DeleteUser(ctx context.Context, idUser int32) error
	GetUserByID(ctx context.Context, idUser int32) (db.User, error)
	UpdateUser(ctx context.Context, params db.UpdateUserParams) error
}

// userRepository é a implementação do UserRepository.
type userRepository struct {
	queries *db.Queries
}

// NewUserRepository cria uma nova instância de UserRepository.
func NewUserRepository(queries *db.Queries) UserRepository {
	return &userRepository{queries: queries}
}

// Implementações dos métodos da interface

// CreateUser cria um novo usuário no banco de dados.
func (r *userRepository) CreateUser(ctx context.Context, params db.CreateUserParams) error {
	return r.queries.CreateUser(ctx, params)
}

// DeleteUser remove um usuário do banco de dados.
func (r *userRepository) DeleteUser(ctx context.Context, idUser int32) error {
	return r.queries.DeleteUser(ctx, idUser)
}

// GetUserByID recupera um usuário pelo ID.
func (r *userRepository) GetUserByID(ctx context.Context, idUser int32) (db.User, error) {
	return r.queries.GetUserByID(ctx, idUser)
}

// UpdateUser atualiza as informações de um usuário no banco de dados.
func (r *userRepository) UpdateUser(ctx context.Context, params db.UpdateUserParams) error {
	return r.queries.UpdateUser(ctx, params)
}

