package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Creamos el router
	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.Run()
}
