package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/tomastropea/meli_bootcamp10/product"
)

func GetProductByIdHandler(products *[]product.Product) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		productIdParam := ctx.Param("id")
		for _, product := range *products {
	
			if strconv.Itoa(product.ID) == productIdParam {
	
				ctx.JSON(200, product)
				return
			}
		}
	
		ctx.String(404, "Product does not exist for given id")
	}
}