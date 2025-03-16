package handlers

import (
	"Fetch/models"
	"Fetch/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"sync"
)

var receiptStore = make(map[string]models.Receipt)

// Outside the current scope for this interview project but this will make sure that there is no race conditions if it where with a real db
var mu sync.RWMutex

func ReceiptSaveHandler(ctx *gin.Context) {
	var receipt models.Receipt

	if err := ctx.ShouldBind(&receipt); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "The receipt is invalid.",
			"details": err.Error(),
		})
		ctx.Abort()
		return
	}

	id := uuid.New().String()

	mu.Lock()
	receiptStore[id] = receipt
	mu.Unlock()

	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func ReceiptPointsHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	mu.RLock()
	receipt, exists := receiptStore[id]
	mu.RUnlock()

	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "No receipt found for that ID.",
		})
		return
	}

	point := services.PointCalc(receipt)
	ctx.JSON(http.StatusOK, gin.H{
		"points": point,
	})
}
