package carrinho

import (
	"cart-api/internal/db"
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Estrutura para receber a requisição de adicionar um item ao carrinho
type CartItemRequest struct {
    CartID    int32   `json:"cart_id" binding:"required"`
    ProductID int32   `json:"product_id" binding:"required"`
    Quantity  int32   `json:"quantity" binding:"required"`
    Price     float64 `json:"price" binding:"required"`
}

// CreateCart cria um novo carrinho para um usuário
func CreateCart(ctx *gin.Context) {
    store := ctx.MustGet("store").(*db.Queries)

    var cartRequest struct {
        UserID int32 `json:"user_id" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&cartRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }

    cartID, err := store.CreateCart(context.Background(), cartRequest.UserID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart", "details": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"cart_id": cartID})
}

// AddItemToCart adiciona um item ao carrinho
func AddItemToCart(ctx *gin.Context) {
    store := ctx.MustGet("store").(*db.Queries)

    var newItem CartItemRequest
    if err := ctx.ShouldBindJSON(&newItem); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }

    priceInt := int32(newItem.Price * 100)

    err := store.AddItemToCart(context.Background(), db.AddItemToCartParams{
        CartID:    newItem.CartID,
        ProductID: newItem.ProductID,
        Quantity:  newItem.Quantity,
        Price:     priceInt,
    })
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart", "details": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

// ListCartItems lista todos os itens de um carrinho
func ListCartItems(ctx *gin.Context) {
    store := ctx.MustGet("store").(*db.Queries)
    cartID, err := strconv.Atoi(ctx.Param("cart_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
        return
    }

    items, err := store.ListCartItems(context.Background(), sql.NullInt32{Int32: int32(cartID), Valid: true})
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
        return
    }

    ctx.JSON(http.StatusOK, items)
}

// TotalCartValue calcula o valor total do carrinho
func TotalCartValue(ctx *gin.Context) {
    store := ctx.MustGet("store").(*db.Queries)
    
    // Validando o cart_id
    cartID, err := strconv.Atoi(ctx.Param("cart_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cart ID must be an integer"})
        return
    }

    log.Printf("Calculating total value for cart ID: %d", cartID)

    // Utilizando o contexto do Gin para a query
    total, err := store.TotalCartValue(ctx.Request.Context(), int32(cartID))
    if err != nil {
        log.Printf("Error calculating total cart value for cart ID %d: %v", cartID, err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total cart value"})
        return
    }

    log.Printf("Total value for cart ID %d: %f", cartID, total)

    // Retornando o total dividido por 100 se for centavos
    ctx.JSON(http.StatusOK, gin.H{"total_value": float64(total) / 100})
}








