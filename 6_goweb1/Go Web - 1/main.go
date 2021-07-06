package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

//estructura producto nombrada con JSON
type Producto struct {
	Id             int     `json:"id"`
	Nombre         string  `json:"nombre"`
	Color          string  `json:"color"`
	Precio         float64 `json:"precio"`
	Stock          int     `json:"stock"`
	Codigo         string  `json:"codigo"`
	Publicado      bool    `json:"publicado"`
	Fecha_creacion string  `json:"fecha_creacion"`
}

func GetAll(c *gin.Context) {
	p := []Producto{
		{Id: 1, Nombre: "Bread Roll", Color: "Puce", Precio: 2.43, Stock: 1, Codigo: "4222-b670", Publicado: true, Fecha_creacion: "19/08/2020"},
		{Id: 2, Nombre: "Ecolab Mild", Color: "Yellow", Precio: 3.73, Stock: 2, Codigo: "65dcc-404", Publicado: false, Fecha_creacion: "08/10/2020"},
	}
	data, err := json.Marshal(p)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	} else {
		c.JSON(200, gin.H{
			"data": string(data),
		})
	}
}

func main() {
	router := gin.Default()
	//ejercicio
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	//ejercicio3
	router.GET("/get-all", GetAll)
	router.Run()
}
