package categoria

import (
	"net/http"
	"strconv"

	services "cart-api/internal/services/categoria"

	"github.com/gin-gonic/gin"
)

// CategoriaHandler define os métodos para os endpoints de categoria.
type CategoriaHandler struct {
	service services.CategoriaService
}

// NewCategoriaHandler cria uma nova instância de CategoriaHandler.
func NewCategoriaHandler(service services.CategoriaService) *CategoriaHandler {
	return &CategoriaHandler{service: service}
}

// request representa o payload para criar ou atualizar uma categoria.
type request struct {
	ID          int32  `json:"id_categoria"` // Alterado para ID maiúsculo para usar em JSON
	Nome        string `json:"nome" binding:"required"`
	Description string `json:"descricao"`
}

// CreateCategory lida com a criação de uma nova categoria.
func (h *CategoriaHandler) CreateCategory(c *gin.Context) {
	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateCategory(c.Request.Context(), req.Nome); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar a categoria: " + err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteCategory lida com a remoção de uma categoria.
func (h *CategoriaHandler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id_categoria") // Alterado para id_categoria
	idCategoria, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.DeleteCategory(c.Request.Context(), int32(idCategoria)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar a categoria: " + err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetCategoryByID recupera uma categoria pelo ID.
func (h *CategoriaHandler) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id_categoria") // Alterado para id_categoria
	idCategoria, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	categoria, err := h.service.GetCategoryByID(c.Request.Context(), int32(idCategoria))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar a categoria: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, categoria)
}

// UpdateCategory atualiza as informações de uma categoria.
func (h *CategoriaHandler) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id_categoria") 
	idCategoria, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ajuste a chamada para passar o contexto correto e os parâmetros corretos
	if err := h.service.UpdateCategory(c.Request.Context(), int32(idCategoria), req.Nome, req.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar a categoria: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Categoria atualizada com sucesso"})
}


