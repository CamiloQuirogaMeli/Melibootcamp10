package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id       string `json: "id"`
	Code     string `json: "code"`
	Currency string `json: "currency"`
	Amount   int    `json: "amount"`
	Emitor   string `json: "emitor"`
	Receptor string `json: "receptor"`
	Date     string `json: "date"`
}

//Prueba con GIN y JSON

func main() {
	fmt.Println("prueba")
	router := gin.Default()
	fmt.Println(router)
	router.GET("/Transactions", getAll)
	router.GET("/query", QueryStr) // algo como query?id=1&amount=10
	router.Run()
}

// para hacer la busqueda de un id hay que escribir http://localhost:8080/query?id=1
func QueryStr(c *gin.Context) {
	id := c.Query("id")
	//code := c.Query("code")
	// currency := c.Query("currency")
	// amount := c.Query("amount")
	// emitor := c.Query("emitor")
	// receptor := c.Query("receptor")
	// date := c.Query("date")

	jsonData, err := ioutil.ReadFile("./Transactions.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	tran := []Transaction{}

	if err := json.Unmarshal(jsonData, &tran); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	aux := []Transaction{}
	for _, t := range tran {
		if id == t.Id {
			aux = append(aux, t)
		}

	}
	c.JSON(200, gin.H{
		"Query": aux,
	})
}

func getAll(c *gin.Context) {

	jsonData, err := ioutil.ReadFile("./Transactions.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	tran := []Transaction{}

	if err := json.Unmarshal(jsonData, &tran); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	c.JSON(200, tran)
	c.JSON(200, gin.H{
		"Transacciones": tran,
	})
}
