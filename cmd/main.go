package main

import (
	"log"
	"net/http"
	"strconv"

	"Transaction-stored/internal/transactions"

	"github.com/gin-gonic/gin"
)

var transactionService *transactions.TransactionService

func main() {
	transactionService = transactions.NewTransactionService()

	r := gin.Default()

	r.POST("/transactionservice/transaction/:id", createTransaction)
	r.GET("/transactionservice/transaction/:id", getTransaction)
	r.GET("/transactionservice/types/:type", getTransactionsByType)
	r.GET("/transactionservice/sum/:id", getTransactionSum)

	if err := r.Run(":8000"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func createTransaction(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var req struct {
		Amount   float64 `json:"amount"`
		Type     string  `json:"type"`
		ParentID *int64  `json:"parent_id,omitempty"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = transactionService.CreateTransaction(id, req.Amount, req.Type, req.ParentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func getTransaction(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := transactionService.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func getTransactionsByType(c *gin.Context) {
	transactionType := c.Param("type")
	ids := transactionService.GetTransactionsByType(transactionType)
	c.JSON(http.StatusOK, ids)
}

func getTransactionSum(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	sum := transactionService.CalculateSum(id)
	c.JSON(http.StatusOK, gin.H{"sum": sum})
}
