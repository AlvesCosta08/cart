package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter configura as rotas da API
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Defina suas rotas aqui
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "up"})
	})

	return router
}


