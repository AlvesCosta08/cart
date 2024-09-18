package carrinho

import (
	"database/sql"
	"net/http"
	"strconv"

	"cart-api/internal/db"

	"github.com/gin-gonic/gin"
)

// CreateCart cria um novo carrinho para um usuário
func CreateCart(ctx *gin.Context) {
	store := ctx.MustGet("store").(*db.Queries)
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	cartID, err := store.CreateCart(ctx, int32(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"cart_id": cartID})
}

// AddItemToCart adiciona um item ao carrinho
func AddItemToCart(ctx *gin.Context) {
	store := ctx.MustGet("store").(*db.Queries)
	var newItem struct {
		CartID    int32   `json:"cart_id" binding:"required"`
		ProductID int32   `json:"product_id" binding:"required"`
		Quantity  int32   `json:"quantity" binding:"required"`
		Price     float64 `json:"price" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&newItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Converte o float64 para int32
	priceInt := int32(newItem.Price * 100) // Supondo que você armazena o preço em centavos

	err := store.AddItemToCart(ctx, db.AddItemToCartParams{
		CartID:    newItem.CartID,
		ProductID: newItem.ProductID,
		Quantity:  newItem.Quantity,
		Price:     priceInt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

// ListCartItems lista todos os itens de um carrinho
func ListCartItems(ctx *gin.Context) {
	store := ctx.MustGet("store").(*db.Queries)
	cartID, err := strconv.Atoi(ctx.Param("cart_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart ID"})
		return
	}

	items, err := store.ListCartItems(ctx, sql.NullInt32{Int32: int32(cartID), Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

// TotalCartValue calcula o valor total do carrinho
func TotalCartValue(ctx *gin.Context) {
	store := ctx.MustGet("store").(*db.Queries)
	cartID, err := strconv.Atoi(ctx.Param("cart_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart ID"})
		return
	}

	total, err := store.TotalCartValue(ctx, sql.NullInt32{Int32: int32(cartID), Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total cart value"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"total_value": total})
}




