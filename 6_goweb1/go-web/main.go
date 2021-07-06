package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id              string  `json:"id"`
	CodeTransaction string  `json:"code_transaction"`
	Currency        string  `json:"currency"`
	Mount           float32 `json:"mount"`
	Sender          string  `json:"sender"`
	Receiver        string  `json:"receiver"`
	TransactionDate string  `json:"transaction_date"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message" : "Hola Alejandro!",
		})
	})
	router.GET("/transactions", func(c *gin.Context){
		getTransactions(c)
	})

	router.Run()
}

func getTransactions(c *gin.Context){
	transactions := []Transaction{}
	jsonData, err := ioutil.ReadFile("./transactions.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err := json.Unmarshal(jsonData, &transactions); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	c.JSON(200, transactions)
}
