package categoria

import (
	"context"

	db "cart-api/db/sqlc"
)

type CategoriaRepository interface {
    CreateCategory(ctx context.Context, params db.CreateCategoryParams) error
    DeleteCategory(ctx context.Context, idCategoria int32) error
    GetCategoryByID(ctx context.Context, idCategoria int32) (db.Categorium, error) 
    UpdateCategory(ctx context.Context, params db.UpdateCategoryParams) error    
}

type categoriaRepository struct {
    queries *db.Queries
}

// NewCategoriaRepository cria uma nova instância de CategoriaRepository.
func NewCategoriaRepository(queries *db.Queries) CategoriaRepository {
    return &categoriaRepository{queries: queries}
}

// CreateCategory cria uma nova categoria no banco de dados.
func (r *categoriaRepository) CreateCategory(ctx context.Context, params db.CreateCategoryParams) error {
    return r.queries.CreateCategory(ctx, params)
}

// DeleteCategory remove uma categoria do banco de dados.
func (r *categoriaRepository) DeleteCategory(ctx context.Context, idCategoria int32) error {
    return r.queries.DeleteCategory(ctx, idCategoria)
}

// GetCategoryByID recupera uma categoria pelo ID.
func (r *categoriaRepository) GetCategoryByID(ctx context.Context, idCategoria int32) (db.Categorium, error) {
    return r.queries.GetCategoryByID(ctx, idCategoria)
}

// UpdateCategory atualiza as informações de uma categoria no banco de dados.
func (r *categoriaRepository) UpdateCategory(ctx context.Context, params db.UpdateCategoryParams) error {
    return r.queries.UpdateCategory(ctx, params)
}
