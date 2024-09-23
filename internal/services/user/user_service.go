package user

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	db "cart-api/db/sqlc"
	userRepository "cart-api/internal/repositories/user"
)

// UserService define os métodos para interagir com usuários.
type UserService interface {
	CreateUser(ctx context.Context, nome, email, senha string) error
	DeleteUser(ctx context.Context, idUser int32) error
	GetUserByID(ctx context.Context, idUser int32) (db.User, error)
	UpdateUser(ctx context.Context, idUser int32, nome, email, senha string) error
}



// userService é a implementação do UserService.
type userService struct {
	repo      userRepository.UserRepository
	validator *validator.Validate
}

// NewUserService cria uma nova instância de UserService.
func NewUserService(repo  userRepository.UserRepository) UserService {
	return &userService{
		repo:      repo,
		validator: validator.New(),
	}
}

// CreateUser cria um novo usuário.
// CreateUser cria um novo usuário.
func (s *userService) CreateUser(ctx context.Context, nome, email, senha string) error {
	var validationErrors []string

	// Validação do nome
	if err := s.validator.Var(nome, "required"); err != nil {
		validationErrors = append(validationErrors, fmt.Sprintf("nome: %v", err))
	}

	// Validação do email
	if err := s.validator.Var(email, "required,email"); err != nil {
		validationErrors = append(validationErrors, fmt.Sprintf("email: %v", err))
	}

	// Validação da senha
	if err := s.validator.Var(senha, "required,min=6"); err != nil {
		validationErrors = append(validationErrors, fmt.Sprintf("senha: %v", err))
	}

	// Se houver erros de validação, retorne-os.
	if len(validationErrors) > 0 {
		return fmt.Errorf("erros de validação: %v", validationErrors)
	}

	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("falha ao gerar hash da senha: %w", err)
	}

	params := db.CreateUserParams{
		Nome:  nome,
		Email: email,
		Senha: string(hashedPassword),
	}

	return s.repo.CreateUser(ctx, params)
}


// DeleteUser remove um usuário.
func (s *userService) DeleteUser(ctx context.Context, idUser int32) error {
	return s.repo.DeleteUser(ctx, idUser)
}

// GetUserByID recupera um usuário pelo ID.
func (s *userService) GetUserByID(ctx context.Context, idUser int32) (db.User, error) {
	return s.repo.GetUserByID(ctx, idUser)
}

// UpdateUser atualiza as informações de um usuário.
func (s *userService) UpdateUser(ctx context.Context, idUser int32, nome, email, senha string) error {
	// Validação
	if err := s.validator.Var(nome, "required"); err != nil {
		return fmt.Errorf("nome: %w", err)
	}
	if err := s.validator.Var(email, "required,email"); err != nil {
		return fmt.Errorf("email: %w", err)
	}

	params := db.UpdateUserParams{
		IDUser: idUser,
		Nome:   nome,
		Email:  email,
	}

	// Se a senha for fornecida, hash e adicionar aos parâmetros.
	if senha != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("falha ao hash da senha: %w", err)
		}
		params.Senha = string(hashedPassword)
	}

	return s.repo.UpdateUser(ctx, params)
}