package main

import (
	"github.com/gin-gonic/gin"
)

var Transaction struct {
	
}

func main() {
	router := gin.Default()

	router.GET("/hello-world", HelloHandler)
	router.GET("/transactions", TransactionHandler)

	router.Run(":8081")
}
