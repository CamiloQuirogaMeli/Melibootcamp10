package main

import (
	f "fmt"

	gin "github.com/gin-gonic/gin"
)

func main() {

	f.Println("Hola mundo")
	router := gin.Default()

	router.GET("/Hello World", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.Run()
}
