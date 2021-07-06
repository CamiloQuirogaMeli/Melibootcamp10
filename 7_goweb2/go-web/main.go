package main

import (
	"github.com/gin-gonic/gin"
)

var Transactions []Transaction

func main() {
	router := gin.Default()
	router.POST("/makeTransaction", Save())
	router.Run()
}
