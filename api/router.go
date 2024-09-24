package api

import (
	"database/sql"
	"net/http"

	"cart-api/api/handlers/categoria"
	handlers "cart-api/api/handlers/user"
	sqlc "cart-api/db/sqlc"
	categoriaRepository "cart-api/internal/repositories/categoria"
	userRepository "cart-api/internal/repositories/user"
	categoriaService "cart-api/internal/services/categoria"
	"cart-api/internal/services/user"

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

	// Instância do repositório e serviço de categorias
	categoriaRepo := categoriaRepository.NewCategoriaRepository(queries)
	categoriaService := categoriaService.NewCategoriaService(categoriaRepo)
	categoriaHandler := categoria.NewCategoriaHandler(categoriaService)

	// Rotas de health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "up"})
	})

	// Rotas para o recurso de usuários
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/", userHandler.GetAllUsersHandler)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// Rotas para o recurso de categorias
	categoriaGroup := router.Group("/categorias")
	{
		categoriaGroup.POST("/", categoriaHandler.CreateCategory)
		categoriaGroup.DELETE("/:id_categoria", categoriaHandler.DeleteCategory)
		categoriaGroup.GET("/:id_categoria", categoriaHandler.GetCategoryByID)
		categoriaGroup.PUT("/:id_categoria", categoriaHandler.UpdateCategory)
	}

	return router
}




