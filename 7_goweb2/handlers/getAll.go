package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/josephcasa/meli_bootcamp10/product"
)

func GetAll(products *[]product.Product) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Products": *products,
		})
	}
}
