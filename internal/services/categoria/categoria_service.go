package service

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"

	db "cart-api/db/sqlc"
	categoriaRepository "cart-api/internal/repositories/categoria"
)

// CategoriaService define os métodos para interagir com categorias.
type CategoriaService interface {
	CreateCategory(ctx context.Context, nome string) error
	DeleteCategory(ctx context.Context, idCategoria int32) error
	GetCategoryByID(ctx context.Context, idCategoria int32) (db.Categorium, error)
	UpdateCategory(ctx context.Context,id int32, nome string,descricao string) error	
}

// categoriaService é a implementação do CategoriaService.
type categoriaService struct {
	repo      categoriaRepository.CategoriaRepository
	validator *validator.Validate
}

// NewCategoriaService cria uma nova instância de CategoriaService.
func NewCategoriaService(repo categoriaRepository.CategoriaRepository) CategoriaService {
	return &categoriaService{
		repo:      repo,
		validator: validator.New(),
	}
}

// CreateCategory cria uma nova categoria.
func (s *categoriaService) CreateCategory(ctx context.Context, nome string) error {
	// Validação do nome
	if err := s.validator.Var(nome, "required"); err != nil {
		return fmt.Errorf("nome: %w", err)
	}

	params := db.CreateCategoryParams{
		Nome: nome,
	}

	return s.repo.CreateCategory(ctx, params)
}

// DeleteCategory remove uma categoria.
func (s *categoriaService) DeleteCategory(ctx context.Context, idCategoria int32) error {
	return s.repo.DeleteCategory(ctx, idCategoria)
}

// GetCategoryByID recupera uma categoria pelo ID.
func (s *categoriaService) GetCategoryByID(ctx context.Context, idCategoria int32) (db.Categorium, error) {
	return s.repo.GetCategoryByID(ctx, idCategoria)
}

// UpdateCategory atualiza as informações de uma categoria.
func (s *categoriaService) UpdateCategory(ctx context.Context,id int32, nome string,descricao string) error {
	// Validação
	if err := s.validator.Var(nome, "required"); err != nil {
		return fmt.Errorf("nome: %w", err)
	}

	params := db.UpdateCategoryParams{
		Nome: nome,
		Description: descricao,
	}

	return s.repo.UpdateCategory(ctx, params)
}


