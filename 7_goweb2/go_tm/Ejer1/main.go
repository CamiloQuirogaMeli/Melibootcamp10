package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transactions struct {
	Transaction []Transaction `json:"transactions"`
}
type Transaction struct {
	Id              int64   `json:"id"`
	TransactionCode string  `json:"transaction_code"`
	CurrencyCode    string  `json:"currency_code"`
	Amount          float64 `json:"amount"`
	Sender          string  `json:"sender"`
	Receiver        string  `json:"receiver"`
}

var allTransactions = []Transaction{}

var maxId int

func main() {

	// Creates a gin router with default middleware
	router := gin.Default()

	//Get all data in the slice
	getAllData()
	maxId = len(allTransactions)
	//First Exercise
	router.GET("/first-route", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola David Santiago Pelaez Parra",
		})
	})

	//Get data Optimized
	router.GET("/transactions", getData())

	//get data filter query
	router.GET("/transactionsFilter", filterData())
	// get data filter params
	router.GET("/transactions/:id", filterId())

	router.POST("/transactions", transactionsPost())

	router.Run() // listen and serve on port 8080

}

func getAllData() {

	jsonFile, err := os.Open("Transactions.json")
	if err != nil {
		fmt.Println("Algo fallo")
	}
	defer jsonFile.Close()
	data := Transactions{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	errJson := json.Unmarshal(byteValue, &data)
	if errJson != nil {
		fmt.Println("Error parseando")
	}
	for i := 0; i < len(data.Transaction); i++ {
		allTransactions = append(allTransactions, data.Transaction[i])
	}
}

func getData() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, allTransactions)
	}
}

func filterData() gin.HandlerFunc {
	return func(c *gin.Context) {
		filtersConst := []string{"Id", "CurrencyCode", "TransactionCode", "Amount", "Sender", "Receiver"}
		filters := c.Request.URL.Query()
		filterReal := []string{}

		for i := 0; i < len(filtersConst); i++ {
			if _, ok := filters[filtersConst[i]]; ok {
				filterReal = append(filterReal, filtersConst[i])

			}
		}
		if len(filterReal) != 0 {
			data := allTransactions
			filtrados := []Transaction{}
			queryAmount, _ := strconv.ParseFloat(c.Query("Amount"), 64)
			queryId, _ := strconv.ParseInt(c.Query("Id"), 10, 64)
			fmt.Println(filterReal)
			for i := 0; i < len(data); i++ {
				if data[i].Id == queryId ||
					data[i].TransactionCode == c.Query("TransactionCode") ||
					data[i].CurrencyCode == c.Query("CurrencyCode") ||
					data[i].Amount == queryAmount ||
					data[i].Sender == c.Query("Sender") ||
					data[i].Receiver == c.Query("Receiver") {
					filtrados = append(filtrados, data[i])
				}
			}
			c.JSON(200, filtrados)
		} else {
			c.String(400, "No funciono el filtrado, intentelo nuevamente")
		}
	}

}

func filterId() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		TransactionFilter := []Transaction{}

		for _, Transaction := range allTransactions {
			if Transaction.Id == paramId {
				TransactionFilter = append(TransactionFilter, Transaction)
				break
			}
		}

		if len(TransactionFilter) > 0 {
			c.JSON(200, TransactionFilter)
		} else {
			c.String(404, "Recurso no encontrado")
		}
	}

}

func transactionsPost() gin.HandlerFunc {

	return func(c *gin.Context) {
		var req Transaction

		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			maxId++
			req.Id = int64(maxId)
			allTransactions = append(allTransactions, req)
			fmt.Println(maxId)
			c.JSON(200, req)
		}

	}
}
