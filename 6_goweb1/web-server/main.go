package main

import (
	f "fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	f.Println("hi")
	router := gin.Default()
	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message" : "/home",
		})
	})

	router.Run()
}