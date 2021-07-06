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

	router.GET("/Transacciones", getDataByQuery)
	router.GET("/Transacciones/:id", getDataByParam)

	router.Run()

}

func getDataByQuery(ctx *gin.Context) {

	transactions := generateTransactionList()
	var filters []Transaction

	for _, t := range transactions {
		if ctx.Query("id") == t.Id {
			filters = append(filters, t)
		}
		if ctx.Query("code") == t.Code {
			filters = append(filters, t)
		}
		if ctx.Query("currency") == t.Currency {
			filters = append(filters, t)
		}
		if ctx.Query("amount") == fmt.Sprintf("%d", t.Amount) {
			filters = append(filters, t)
		}
		if ctx.Query("emitor") == t.Emitor {
			filters = append(filters, t)
		}
		if ctx.Query("r©eceptor") == t.Receptor {
			filters = append(filters, t)
		}
		if ctx.Query("date") == t.Date {
			filters = append(filters, t)
		}
	}
	if len(filters) == 0 {
		ctx.String(400, "No hay resultados!!")
	} else {
		ctx.JSON(200, filters)
	}

}

func getDataByParam(ctx *gin.Context) {

	transactions := generateTransactionList()
	var filters []Transaction

	for _, t := range transactions {
		if ctx.Param("id") == t.Id {
			filters = append(filters, t)
		}
	}

	if len(filters) == 0 {
		ctx.String(400, "No hay resultados!!")
	} else {
		ctx.JSON(200, filters)
	}

}
func generateTransactionList() []Transaction {
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

	return t

}
