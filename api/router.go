package api

import (
	sqlc "cart-api/db/sqlc"

	handlers "cart-api/api/handlers/user"
	userRepository "cart-api/internal/repositories/user"
	"cart-api/internal/services/user"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura as rotas da API
func SetupRouter(dbConn *sql.DB) *gin.Engine {
	router := gin.Default()

	// Instância do repositório e serviço de usuários
	queries := sqlc.New(dbConn)
	userRepo := userRepository.NewUserRepository(queries)
	userService := user.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Rotas de health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "up"})
	})

	// Rotas para o recurso de usuários
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)      // Criar usuário
		userRoutes.GET("/:id", userHandler.GetUserByID)   // Obter usuário por ID
		userRoutes.PUT("/:id", userHandler.UpdateUser)    // Atualizar usuário por ID
		userRoutes.DELETE("/:id", userHandler.DeleteUser) // Deletar usuário por ID
	}

	return router
}



