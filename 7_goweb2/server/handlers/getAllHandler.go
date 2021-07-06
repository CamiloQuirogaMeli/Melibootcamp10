package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tomastropea/meli_bootcamp10/product"
)

func GetAllHandler(products *[]product.Product) gin.HandlerFunc {

	return func(ctx *gin.Context) {
	
		ctx.JSON(200, gin.H{
			"product.products": *products,
		}) 
	}
}