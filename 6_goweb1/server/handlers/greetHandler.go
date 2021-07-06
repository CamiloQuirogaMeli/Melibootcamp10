package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GreetUserHandler(ctx *gin.Context) {
	
	ctx.JSON(200, gin.H{
		"message": fmt.Sprintf("Hola %s", NAME),
	})
}
