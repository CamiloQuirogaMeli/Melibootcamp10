package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Productos struct {
	Id             int
	Nombre         string
	Color          string
	Precio         float64
	Stock          int
	Codigo         string
	Publicado      bool
	Fecha_creacion string
}



func main() {

	// Creates a gin router with default middleware
	router := gin.Default()

	// A handler for GET request on /example
	router.GET("/hello-camilo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Camilo!",
		}) // gin.H is a shortcut for map[string]interface{}
	})

	router.GET("/productos", func(c *gin.Context) {
		dat, err := ioutil.ReadFile("../Ejercicio1/Productos.json")
		if err != nil {
			c.JSON(500, gin.H{
				"error": "No se pudo leer el archivo",
			})
		} else {
			var producto []Productos
			json.Unmarshal(dat, &producto)
			c.String(200, "%v", producto)
		}

	})
	router.Run() // listen and serve on port 8080
}
