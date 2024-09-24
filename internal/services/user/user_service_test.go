package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	db "cart-api/db/sqlc"
)

// MockUserRepository é uma estrutura que simula o comportamento do UserRepository.
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, params db.CreateUserParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, idUser int32) error {
	args := m.Called(ctx, idUser)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, idUser int32) (db.User, error) {
	args := m.Called(ctx, idUser)
	return args.Get(0).(db.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, params db.UpdateUserParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

// Adiciona o método GetAllUsers no MockUserRepository.
func (m *MockUserRepository) GetAllUsers(ctx context.Context) ([]db.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]db.User), args.Error(1)
}

// TestCreateUserSuccess testa a criação de um usuário com sucesso.
func TestCreateUserSuccess(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	nome := "John Doe"
	email := "johndoe@example.com"
	senha := "password123"

	// Simula o comportamento esperado do repositório ao criar um usuário.
	mockRepo.On("CreateUser", ctx, mock.Anything).Return(nil)

	err := service.CreateUser(ctx, nome, email, senha)

	// Verifica se não houve erros e se o método foi chamado corretamente.
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "CreateUser", ctx, mock.Anything)
}

// TestCreateUserValidationError testa erros de validação durante a criação de um usuário.
func TestCreateUserValidationError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	nome := ""          // Nome vazio
	email := "invalid"  // E-mail inválido
	senha := "12"       // Senha curta

	err := service.CreateUser(ctx, nome, email, senha)

	// Verifica se os erros de validação são capturados corretamente.
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nome")
	assert.Contains(t, err.Error(), "email")
	assert.Contains(t, err.Error(), "senha") // Agora deve incluir "senha"

	// Verifica se o método CreateUser do repositório não foi chamado.
	mockRepo.AssertNotCalled(t, "CreateUser", mock.Anything, mock.Anything)
}

// TestCreateUserRepositoryError testa a criação de um usuário quando ocorre um erro no repositório.
func TestCreateUserRepositoryError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	nome := "John Doe"
	email := "johndoe@example.com"
	senha := "password123"

	// Simula um erro no repositório.
	mockRepo.On("CreateUser", ctx, mock.Anything).Return(errors.New("database error"))

	err := service.CreateUser(ctx, nome, email, senha)

	// Verifica se o erro é retornado corretamente.
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database error")
	mockRepo.AssertCalled(t, "CreateUser", ctx, mock.Anything)
}

// TestDeleteUserSuccess testa a exclusão de um usuário com sucesso.
func TestDeleteUserSuccess(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	idUser := int32(1)

	// Simula o comportamento esperado do repositório ao deletar um usuário.
	mockRepo.On("DeleteUser", ctx, idUser).Return(nil)

	err := service.DeleteUser(ctx, idUser)

	// Verifica se não houve erros e se o método foi chamado corretamente.
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteUser", ctx, idUser)
}

// TestDeleteUserNotFound testa a tentativa de exclusão de um usuário que não existe.
func TestDeleteUserNotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	idUser := int32(999) // Supondo que o ID 999 não exista.

	// Simula um erro ao tentar deletar um usuário que não existe.
	mockRepo.On("DeleteUser", ctx, idUser).Return(errors.New("user not found"))

	err := service.DeleteUser(ctx, idUser)

	// Verifica se o erro é retornado corretamente.
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")
	mockRepo.AssertCalled(t, "DeleteUser", ctx, idUser)
}

// TestGetUserByIDSuccess testa a recuperação de um usuário com sucesso.
func TestGetUserByIDSuccess(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	idUser := int32(1)
	expectedUser := db.User{IDUser: idUser, Nome: "John Doe", Email: "johndoe@example.com"}

	// Simula o comportamento esperado do repositório ao recuperar um usuário.
	mockRepo.On("GetUserByID", ctx, idUser).Return(expectedUser, nil)

	user, err := service.GetUserByID(ctx, idUser)

	// Verifica se o usuário foi retornado corretamente e sem erros.
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertCalled(t, "GetUserByID", ctx, idUser)
}

// TestGetUserByIDNotFound testa a recuperação de um usuário que não existe.
func TestGetUserByIDNotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	idUser := int32(999) // Supondo que o ID 999 não exista.

	// Simula um erro ao tentar recuperar um usuário que não existe.
	mockRepo.On("GetUserByID", ctx, idUser).Return(db.User{}, errors.New("user not found"))

	user, err := service.GetUserByID(ctx, idUser)

	// Verifica se o erro é retornado corretamente e se o usuário retornado está vazio.
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")
	assert.Equal(t, db.User{}, user)
	mockRepo.AssertCalled(t, "GetUserByID", ctx, idUser)
}

// TestUpdateUserSuccess testa a atualização de um usuário com sucesso.
func TestUpdateUserSuccess(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	idUser := int32(1)
	nome := "John Doe Updated"
	email := "johnupdated@example.com"
	senha := "newpassword"

	// Simula o comportamento esperado do repositório ao atualizar um usuário.
	mockRepo.On("UpdateUser", ctx, mock.Anything).Return(nil)

	err := service.UpdateUser(ctx, idUser, nome, email, senha)

	// Verifica se não houve erros e se o método foi chamado corretamente.
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "UpdateUser", ctx, mock.Anything)
}

// TestUpdateUserNoPassword testa a atualização de um usuário sem alterar a senha.
func TestUpdateUserNoPassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	idUser := int32(1)
	nome := "John Doe Updated"
	email := "johnupdated@example.com"
	senha := ""

	// Simula o comportamento esperado do repositório ao atualizar um usuário.
	mockRepo.On("UpdateUser", ctx, mock.Anything).Return(nil)

	err := service.UpdateUser(ctx, idUser, nome, email, senha)

	// Verifica se não houve erros e se o método foi chamado corretamente.
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "UpdateUser", ctx, mock.Anything)
}

// TestUpdateUserRepositoryError testa a atualização de um usuário quando ocorre um erro no repositório.
func TestUpdateUserRepositoryError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	idUser := int32(1)
	nome := "John Doe Updated"
	email := "johnupdated@example.com"
	senha := "newpassword"

	// Simula um erro no repositório ao atualizar um usuário.
	mockRepo.On("UpdateUser", ctx, mock.Anything).Return(errors.New("update error"))

	err := service.UpdateUser(ctx, idUser, nome, email, senha)

	// Verifica se o erro é retornado corretamente.
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "update error")
	mockRepo.AssertCalled(t, "UpdateUser", ctx, mock.Anything)
}

// TestGetAllUsersSuccess testa a recuperação de todos os usuários com sucesso.
func TestGetAllUsersSuccess(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	expectedUsers := []db.User{
		{IDUser: 1, Nome: "John Doe", Email: "johndoe@example.com"},
		{IDUser: 2, Nome: "Jane Doe", Email: "janedoe@example.com"},
	}

	// Simula o comportamento esperado do repositório ao recuperar todos os usuários.
	mockRepo.On("GetAllUsers", ctx).Return(expectedUsers, nil)

	users, err := service.GetAllUsers(ctx)

	// Verifica se os usuários foram retornados corretamente e sem erros.
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
	mockRepo.AssertCalled(t, "GetAllUsers", ctx)
}

// TestGetAllUsersError testa a recuperação de todos os usuários quando ocorre um erro no repositório.
func TestGetAllUsersError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()

	// Simula um erro no repositório ao tentar obter todos os usuários.
	mockRepo.On("GetAllUsers", ctx).Return([]db.User{}, errors.New("database error"))

	users, err := service.GetAllUsers(ctx)

	// Verifica se o erro é retornado corretamente.
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database error")
	assert.Empty(t, users) // Verifica se a lista de usuários está vazia.
	mockRepo.AssertCalled(t, "GetAllUsers", ctx)
}

// TestCreateUserEmailExists testa a criação de um usuário com e-mail já existente.
func TestCreateUserEmailExists(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	nome := "John Doe"
	email := "johndoe@example.com"
	senha := "password123"

	// Simula um erro de e-mail já existente.
	mockRepo.On("CreateUser", ctx, mock.Anything).Return(errors.New("email already exists"))

	err := service.CreateUser(ctx, nome, email, senha)

	// Verifica se o erro é retornado corretamente.
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email already exists")
	mockRepo.AssertCalled(t, "CreateUser", ctx, mock.Anything)
}

// TestUpdateUserEmailExists testa a atualização de um usuário com e-mail já existente.
func TestUpdateUserEmailExists(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	idUser := int32(1)
	nome := "John Doe Updated"
	email := "johnupdated@example.com"
	senha := "newpassword"

	// Simula um erro ao tentar atualizar para um e-mail que já existe.
	mockRepo.On("UpdateUser", ctx, mock.Anything).Return(errors.New("email already exists"))

	err := service.UpdateUser(ctx, idUser, nome, email, senha)

	// Verifica se o erro é retornado corretamente.
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email already exists")
	mockRepo.AssertCalled(t, "UpdateUser", ctx, mock.Anything)
}

// TestDeleteUserAlreadyDeleted testa a exclusão de um usuário que já foi deletado.
func TestDeleteUserAlreadyDeleted(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	idUser := int32(1)

	// Simula um erro ao tentar deletar um usuário que não existe mais.
	mockRepo.On("DeleteUser", ctx, idUser).Return(errors.New("user already deleted"))

	err := service.DeleteUser(ctx, idUser)

	// Verifica se o erro é retornado corretamente.
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user already deleted")
	mockRepo.AssertCalled(t, "DeleteUser", ctx, idUser)
}

// TestGetAllUsersEmpty testa a recuperação de todos os usuários quando não há nenhum usuário.
func TestGetAllUsersEmpty(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()

	// Simula a recuperação de uma lista vazia de usuários.
	mockRepo.On("GetAllUsers", ctx).Return([]db.User{}, nil)

	users, err := service.GetAllUsers(ctx)

	// Verifica se não houve erros e que a lista de usuários está vazia.
	assert.NoError(t, err)
	assert.Empty(t, users)
	mockRepo.AssertCalled(t, "GetAllUsers", ctx)
}