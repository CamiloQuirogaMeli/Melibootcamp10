package handlers

import (
	f "fmt"

	"github.com/gin-gonic/gin"
)

const Name = "Joseph Caicedo"

func Greet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": f.Sprint("Hello ", Name),
	})
}
