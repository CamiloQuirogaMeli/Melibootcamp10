package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const PORT int = 3500

func main() {
	fmt.Println("Hola mundo")

	router := gin.Default()
	fmt.Println(router)

	router.GET("/HelloWorld", sendJSON)

	router.Run()

}

func sendJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World2",
	})
}
