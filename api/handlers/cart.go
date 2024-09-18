package handlers

import (
	"cart-api/internal/db"
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateCart(c *gin.Context, store *db.Queries) {
    userIDStr := c.PostForm("user_id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    id, err := store.CreateCart(context.Background(), int32(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"cart_id": id})
}

func AddItemToCart(c *gin.Context, store *db.Queries) {
    var params db.AddItemToCartParams

    if err := c.BindJSON(&params); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    err := store.AddItemToCart(context.Background(), params)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func ListCartItems(c *gin.Context, store *db.Queries) {
    cartIDStr := c.Param("cart_id")
    cartID, err := strconv.Atoi(cartIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
        return
    }

    items, err := store.ListCartItems(context.Background(), sql.NullInt32{Int32: int32(cartID), Valid: true})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list cart items"})
        return
    }

    c.JSON(http.StatusOK, items)
}

func GetCartTotals(c *gin.Context, store *db.Queries) {
    cartIDStr := c.Param("cart_id")
    cartID, err := strconv.Atoi(cartIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
        return
    }

    totalItems, err := store.TotalCartItems(context.Background(), sql.NullInt32{Int32: int32(cartID), Valid: true})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total items"})
        return
    }

    totalValue, err := store.TotalCartValue(context.Background(), sql.NullInt32{Int32: int32(cartID), Valid: true})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total value"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "total_items": totalItems,
        "total_value": totalValue,
    })
}
