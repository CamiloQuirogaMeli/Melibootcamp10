package main

import "github.com/gin-gonic/gin"

func main() {
	// Creates a gin router with default middleware
	router := gin.Default()

	// A handler for GET request on /hello-world
	router.GET("/first-route", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hola David Santiago Pelaez Parra",
		}) // gin.H is a shortcut for map[string]interface{}
	})

	router.Run() // listen and serve on port 8080

}
