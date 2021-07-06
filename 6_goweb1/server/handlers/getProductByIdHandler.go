package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetProductByIdHandler(ctx *gin.Context) {

	products, err := LoadProductsFromFile(PRODUCTS_PATH)
	
	if err != nil {

		ctx.String(500, err.Error())
		return
	}

	productIdParam := ctx.Param("id")
	for _, product := range products {

		if strconv.Itoa(product.ID) == productIdParam {

			ctx.JSON(200, product)
			return
		}
	}

	ctx.String(404, "Product does not exist for given id")
}