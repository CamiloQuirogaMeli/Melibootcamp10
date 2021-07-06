package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id             int
	Nombre         string
	Color          string
	Precio         float64
	Stock          int
	Codigo         string
	Publicado      bool
	Fecha_creacion string
}

func saludo(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hola Valeria",
	})
}

func getAll(ctx *gin.Context) {
	data, err := ioutil.ReadFile("product.json")

	if err != nil {
		log.Fatal(err)
	}
	var p []Producto
	if err := json.Unmarshal([]byte(data), &p); err != nil {
		log.Fatal(err)
	}
	ctx.JSON(200, gin.H{
		"Productos": p,
	})
}

func main() {
	router := gin.Default()

	router.GET("/hello-world", saludo)
	router.GET("/productos", getAll)

	router.Run()
}
