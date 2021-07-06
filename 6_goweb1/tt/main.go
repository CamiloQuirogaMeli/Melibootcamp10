package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id             int     `json:"id"`
	CodTransaccion string  `json:"codTransaccion"`
	Moneda         string  `json:"moneda"`
	Monto          float64 `json:"monto"`
	Emisor         int     `json:"emisor"`
	Receptor       int     `json:"receptor"`
	Fecha          string  `json:"fecha"`
}

type Handler struct{}

func main() {

	router := gin.Default()

	h := &Handler{}

	router.GET("/ej2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Luciano Panizza",
		})
	})

	router.GET("/transactions", h.GetAll)
	router.GET("/transactions/:id", h.SearchTransaction)

	router.Run()

}

func (h *Handler) GetAll(c *gin.Context) {
	jsonData, err := ioutil.ReadFile("./transactions.json")
	transacciones := []Transaction{}
	cantidadFriltados := 0
	var filtrados []Transaction
	if err != nil {
		fmt.Printf("Erro leyendo el json")
	}

	if err := json.Unmarshal(jsonData, &transacciones); err != nil {
		c.JSON(500, gin.H{
			"error": "No se pudo deserializar el archivo.",
		})
	}

	for _, transaccion := range transacciones {
		if c.Query("moneda") == transaccion.Moneda {
			filtrados = append(filtrados, transaccion)
			cantidadFriltados++
		}
	}

	if cantidadFriltados > 0 {
		c.String(200, "%v", filtrados)
	} else {
		c.String(404, "No se encontraron coincidencias")
	}

}

func (h *Handler) SearchTransaction(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("Error parseando el ID")
	}

	jsonData, err := ioutil.ReadFile("./transactions.json")
	transacciones := []Transaction{}

	if err != nil {
		fmt.Printf("Error leyendo el json")
	}

	if err := json.Unmarshal(jsonData, &transacciones); err != nil {
		c.JSON(500, gin.H{
			"error": "No se pudo deserializar el archivo.",
		})
	}

	for _, transaccion := range transacciones {
		if id == transaccion.Id {
			c.String(200, "%v", transaccion)
			return
		}
	}
	c.String(404, "No se encontraron coincidencias")

}
