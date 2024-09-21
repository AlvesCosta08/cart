package main

import (
	"cart-api/api/handlers/carrinho"
	"cart-api/api/handlers/produtos"
	"cart-api/internal/db"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Conectar ao banco de dados PostgreSQL
	dbConn, err := sql.Open("postgres", "user=postgres password=postgres dbname=shopping_cart sslmode=disable") // Ajuste a DSN conforme necessário
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer dbConn.Close()

	// Inicializar o SQLC Queries
	queries := db.New(dbConn)

	// Inicializar o router Gin
	r := gin.Default()

	// Middleware para adicionar o store ao contexto
	r.Use(func(c *gin.Context) {
		c.Set("store", queries)
		c.Next()
	})

	// Rotas para produtos
	r.GET("/products", produtos.ListProducts)
	r.GET("/products/:id", produtos.GetProduct)
	r.POST("/products", produtos.CreateProduct)

	// Rotas para carrinho
	r.POST("/carts", carrinho.CreateCart)
	r.POST("/carts/items", carrinho.AddItemToCart)
	r.GET("/carts/:cart_id/items", carrinho.ListCartItems)
	r.GET("/carts/:cart_id/total", carrinho.TotalCartValue)
	r.POST("/carts/:cart_id/checkout", carrinho.CheckoutCart)


	// Iniciar o servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
