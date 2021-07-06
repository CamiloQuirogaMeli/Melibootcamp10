package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type Transaction struct {
	Id int `json:"id"`
	TransactionCode string `json:"transaction_code"`
	Money int `json:"money"`
	Emitter string `json:"emitter"`
	Receiver string `json:"receiver"`
	Date string `json:"date"`
}

func GetAll(ctx *gin.Context) {
	content, err := ioutil.ReadFile("./go-web/transactions.json")
	if err != nil {
		log.Fatal(err)
	}
	println(string(content))
	transactions:= []Transaction{}
	if errJ := json.Unmarshal(content, &transactions); errJ != nil {
		log.Fatal(err)
	}
	ctx.JSON(200,gin.H{
		"Message": transactions,
	})
}

func main() {
	router := gin.Default()

	router.GET("/transactions", GetAll)
	router.Run()
}
