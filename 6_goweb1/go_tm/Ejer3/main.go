package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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

	router.Run() // listen and serve on port 8080

}

func getAll() []byte {
	jsonFile, err := os.Open("Transactions.json")
	if err != nil {
		fmt.Println("Algo fallo")
	}
	defer jsonFile.Close()
	// _ = json.Unmarshal([]byte(file), &data)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}
