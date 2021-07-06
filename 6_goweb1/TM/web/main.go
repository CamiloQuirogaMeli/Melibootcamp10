package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id       string `json:"id"`
	Code     string `json:"code"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
	Emitor   string `json:"emitor"`
	Receptor string `json:"receptor"`
	Date     string `json:"date"`
}

const PORT int = 3500

func main() {

	router := gin.Default()

	router.GET("/Transacciones", getAll)
	router.GET("/Hola", getName)

	router.Run()

}

func getName(c *gin.Context) {

	c.JSON(200, gin.H{"message": "Hola Santiago"})
}

func getAll(c *gin.Context) {
	jsonData, err := ioutil.ReadFile("./transactions.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	t := []Transaction{}

	if err := json.Unmarshal(jsonData, &t); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	c.JSON(200, t)
}
