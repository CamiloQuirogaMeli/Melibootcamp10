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
	Id       int    `json: "id"`
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

	//endpoint,						esta funciones es handler
	// router.GET("/Helloworld", func(c *gin.Context) {

	// 	c.JSON(200, gin.H{
	// 		"message": "Hello Martin",
	// 	})
	// })
	router.Run()

	// router.GET("/tematicas", func(c *gin.Context) {

	// 	c.JSON(200, gin.H{
	// 		"message": "Hello Martin",
	// 	})
	// })
	// router.Run()

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
		"Transacciones": tran
	})
}
