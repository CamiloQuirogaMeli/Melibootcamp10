package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("prueba")
	router := gin.Default()

	router.GET("/Helloworld", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
	router.Run()
}
