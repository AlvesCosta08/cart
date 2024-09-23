package user

import (
	"context"
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
	senha := "123"      // Senha curta

	err := service.CreateUser(ctx, nome, email, senha)

	// Verifica se os erros de validação são capturados corretamente.
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nome")
	assert.Contains(t, err.Error(), "email")
	assert.Contains(t, err.Error(), "senha")

	// Verifica se o método CreateUser do repositório não foi chamado.
	mockRepo.AssertNotCalled(t, "CreateUser", mock.Anything, mock.Anything)
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

