package user

import (
	"net/http"
	"strconv"

	"cart-api/internal/services/user" // ajuste conforme o caminho correto do pacote de serviços

	"github.com/gin-gonic/gin"
)

// UserHandler define os métodos para gerenciar as requisições de usuários.
type UserHandler struct {
	service user.UserService
}

// NewUserHandler cria uma nova instância de UserHandler.
func NewUserHandler(service user.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser lida com a criação de um novo usuário.
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input struct {
		Nome  string `json:"nome" binding:"required"`
		Email string `json:"email" binding:"required,email"`
		Senha string `json:"senha" binding:"required,min=3"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateUser(c, input.Nome, input.Email, input.Senha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
}

// GetUserByID lida com a recuperação de um usuário pelo ID.
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id") // ID é uma string

	// Converte o ID de string para int32
	idUser, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := h.service.GetUserByID(c, int32(idUser))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser lida com a atualização das informações de um usuário.
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id") // ID é uma string

	// Converte o ID de string para int32
	idUser, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		Nome  string `json:"nome" binding:"required"`
		Email string `json:"email" binding:"required,email"`
		Senha string `json:"senha,omitempty"` 
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.UpdateUser(c, int32(idUser), input.Nome, input.Email, input.Senha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso"})
}

// DeleteUser lida com a exclusão de um usuário.
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id") // ID é uma string

	// Converte o ID de string para int32
	idUser, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.service.DeleteUser(c, int32(idUser))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil) 
}

// GetAllUsersHandler retorna uma lista de todos os usuários.
func (h *UserHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := h.service.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"})
		return
	}
	c.JSON(http.StatusOK, users)
}