package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//EJERCICIO 2 MOSTRAR MENSAJE CON NOMBRE

	router := gin.Default()

	// A handler for GET request on /example

	router.GET("/saludo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Martina!",
		}) // gin.H is a shortcut for map[string]interface{}
	})

	router.Run()
}
