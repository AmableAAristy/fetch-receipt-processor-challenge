package main

import (
	"Fetch/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.POST("/receipts/process", handlers.ReceiptSaveHandler)
	r.GET("/receipts/:id/points", handlers.ReceiptPointsHandler)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
