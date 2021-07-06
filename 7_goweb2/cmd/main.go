package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Transaction struct {
	ID         int     `json:"id,omitempty"`
	Code       string  `json:"code_transaction,omitempty"`
	Currency   string  `json:"currency,omitempty"`
	Mount      float64 `json:"mount,omitempty"`
	Dispatcher string  `json:"dispatcher,omitempty"`
	Receiver   string  `json:"receiver,omitempty"`
	Date       string  `json:"transaction_date,omitempty"`
}

func generateID(transactions []Transaction) int {
	if len(transactions) == 0 {
		return 1
	}
	newID := len(transactions) + 1
	return newID
}

const validToken = "1234"

func main() {
	router := gin.Default()
	var transactions []Transaction
	// A handler for GET request on /hello-world
	router.GET("/transactions", func(c *gin.Context) {
		c.JSON(http.StatusOK, transactions)
	})

	router.POST("/transactions", validateToken(), validateParams(), func(c *gin.Context) {
		var t Transaction
		var id int
		if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id = generateID(transactions)
		t.ID = id
		transactions = append(transactions, t)
		c.Status(http.StatusOK)
	})

	router.Run()
}

func validateParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		var t Transaction
		if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if t.Code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo code_transaction es requerido"})
			c.Abort()
		}
		if t.Currency == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo currency es requerido"})
			c.Abort()
		}
		if t.Date == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo transaction_date es requerido"})
			c.Abort()
		}
		if t.Dispatcher == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo dispatcher es requerido"})
			c.Abort()
		}
		if t.Receiver == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo receiver es requerido"})
			c.Abort()
		}
		if t.Mount == 0.0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El campo mount es requerido"})
			c.Abort()
		}
	}
}

func validateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != validToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No tiene permisos para realizar la peticion solicitada"})
		}
		c.Abort()
	}
}
