package handlers

import (
	"cart-api/internal/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListProducts(ctx *gin.Context, store *db.Queries) {
    products, err := store.ListProducts(ctx)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
        return
    }
    ctx.JSON(http.StatusOK, products)
}

func GetProduct(ctx *gin.Context, store *db.Queries) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
        return
    }

    product, err := store.GetProductById(ctx, int32(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
        return
    }

    ctx.JSON(http.StatusOK, product)
}

func CreateProduct(ctx *gin.Context, store *db.Queries) {
    var newProduct struct {
        Name  string `json:"name" binding:"required"`
        Price string `json:"price" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&newProduct); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    price, err := strconv.ParseFloat(newProduct.Price, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price format"})
        return
    }

    // Converte o float64 para string
    priceStr := strconv.FormatFloat(price, 'f', 2, 64)

    createdProduct, err := store.CreateProduct(ctx, db.CreateProductParams{
        Name:  newProduct.Name,
        Price: priceStr,
    })
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }

    ctx.JSON(http.StatusCreated, createdProduct)
}
