package main

import (
	"fmt"
	"net/http"

	"github.com/MarcosZabala/go-web/datastore"
	"github.com/gin-gonic/gin"
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

const jsonPath = "../transaction.json"

func main() {
	router := gin.Default()
	ds := datastore.NewDS(jsonPath)
	// A handler for GET request on /hello-world
	router.GET("/hello/:name", func(c *gin.Context) {
		name, ok := c.Params.Get("name")
		if !ok {
			c.JSON(404, gin.H{})
		}
		c.JSON(200, gin.H{
			"message": "Hello " + name,
		})
	})

	router.GET("/transactions", func(c *gin.Context) {
		queryParams := c.Request.URL.Query()
		fmt.Println(len(queryParams))
		id := c.Query("id")
		code := c.Query("code_transaction")
		currency := c.Query("currency")
		mount := c.Query("mount")
		dispatcher := c.Query("dispatcher")
		receiver := c.Query("receiver")
		date := c.Query("transaction_date")
		transactions := ds.GetByFilter(id, code, currency, mount, dispatcher, receiver, date)
		c.JSON(http.StatusOK, transactions)
	})

	router.GET("/transactions/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		t := ds.GetByID(id)
		if t == (datastore.Transaction{}) {
			c.JSON(http.StatusNotFound, t)
			return
		}
		c.JSON(http.StatusOK, t)
	})

	router.Run()
}
