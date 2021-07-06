package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

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

func main() {
	// Creates a gin router with default middleware
	router := gin.Default()

	// A handler for GET request on /hello-world
	router.GET("/first-route", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hola David Santiago Pelaez Parra",
		}) // gin.H is a shortcut for map[string]interface{}
	})

	//Handler for Get request Transactions
	router.GET("/transactions", func(c *gin.Context) {
		data := Transactions{}
		jsonData := getAll()
		errJson := json.Unmarshal(jsonData, &data)
		if errJson != nil {
			fmt.Println(errJson)
			fmt.Println("Fallo algo")
		}
		c.JSON(200, data)
	})
	router.GET("/transactionsFilter", filterData)

	router.GET("/transactions/:id", filterId)

	router.Run() // listen and serve on port 8080

}

func getAll() []byte {
	jsonFile, err := os.Open("Transactions.json")
	if err != nil {
		fmt.Println("Algo fallo")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func filterData(c *gin.Context) {
	filtersConst := []string{"Id", "CurrencyCode", "TransactionCode", "Amount", "Sender", "Receiver"}
	filters := c.Request.URL.Query()
	filterReal := []string{}
	values := []string{}
	for i := 0; i < len(filtersConst); i++ {
		if val, ok := filters[filtersConst[i]]; ok {
			filterReal = append(filterReal, filtersConst[i])
			values = append(values, strings.Join(val, ""))
		}
	}
	if len(filterReal) != 0 {
		data := Transactions{}
		jsonData := getAll()
		errJson := json.Unmarshal(jsonData, &data)
		if errJson != nil {
			fmt.Println(errJson)
			fmt.Println("Fallo algo")
		}
		var filtrados []Transaction
		queryAmount, _ := strconv.ParseFloat(c.Query("Amount"), 64)
		queryId, _ := strconv.ParseInt(c.Query("Id"), 10, 64)
		for i := 0; i < len(data.Transaction); i++ {
			if data.Transaction[i].Id == queryId ||
				data.Transaction[i].TransactionCode == c.Query("TransactionCode") ||
				data.Transaction[i].CurrencyCode == c.Query("CurrencyCode") ||
				data.Transaction[i].Amount == queryAmount ||
				data.Transaction[i].Sender == c.Query("Sender") ||
				data.Transaction[i].Receiver == c.Query("Receiver") {
				filtrados = append(filtrados, data.Transaction[i])
			}
		}
		c.JSON(200, filtrados)
	} else {
		c.String(400, "No funciono el filtrado, intentelo nuevamente")
	}

}

func filterId(c *gin.Context) {
	paramId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	data := Transactions{}
	jsonData := getAll()
	errJson := json.Unmarshal(jsonData, &data)
	if errJson != nil {
		fmt.Println(errJson)
		fmt.Println("Fallo algo")
	}
	TransactionFilter := []Transaction{}

	for _, Transaction := range data.Transaction {
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
