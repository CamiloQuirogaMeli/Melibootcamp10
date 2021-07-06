package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id             int     `json:"id"`
	CodTransaccion string  `json:"codTransaccion" binding:"required"`
	Moneda         string  `json:"moneda" binding:"required"`
	Monto          float64 `json:"monto" binding:"required"`
	Emisor         int     `json:"emisor" binding:"required"`
	Receptor       int     `json:"receptor" binding:"required"`
	Fecha          string  `json:"fecha" binding:"required"`
}

type Handler struct{}

func main() {

	router := gin.Default()
	db := []*Transaction{}

	h := &Handler{}

	router.GET("/ej2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Luciano Panizza",
		})
	})

	router.GET("/transactions", h.GetAll)
	router.GET("/transactions/:id", h.SearchTransaction)

	router.POST("/transactions", GuardarReq(db))

	router.Run()

}

func GuardarReq(db []*Transaction) gin.HandlerFunc {

	return func(c *gin.Context) {
		var transaccionReq Transaction

		token := c.Request.Header.Get("token")

		if token != "ASD321" {
			c.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}

		if err := c.ShouldBindJSON(&transaccionReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		transaccionReq.Id = len(db)

		db = append(db, &transaccionReq)

		c.String(200, "Se grego la transaccion %v", transaccionReq)

	}
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
