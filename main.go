package main

import (
	"cart-api/api/handlers"
	"cart-api/internal/db"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
    conn, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/shopping_cart?sslmode=disable")
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }

    store := db.New(conn)
    router := gin.Default()

    // Rotas de produto
    router.GET("/api/products", func(ctx *gin.Context) {
        handlers.ListProducts(ctx, store)
    })

    router.GET("/api/products/:id", func(ctx *gin.Context) {
        handlers.GetProduct(ctx, store)
    })

    router.POST("/api/products", func(ctx *gin.Context) {
        handlers.CreateProduct(ctx, store)
    })

    // Rotas do carrinho
    router.POST("/api/cart", func(ctx *gin.Context) {
        handlers.CreateCart(ctx, store)
    })

    router.POST("/api/cart/items", func(ctx *gin.Context) {
        handlers.AddItemToCart(ctx, store)
    })

    router.GET("/api/cart/items/:cart_id", func(ctx *gin.Context) {
        handlers.ListCartItems(ctx, store)
    })

    router.GET("/api/cart/totals/:cart_id", func(ctx *gin.Context) {
        handlers.GetCartTotals(ctx, store)
    })

    router.Run(":8080")
}

