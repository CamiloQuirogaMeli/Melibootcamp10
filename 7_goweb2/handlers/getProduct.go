package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/josephcasa/meli_bootcamp10/product"
)

func GetProduct(products *[]product.Product) gin.HandlerFunc {

	return func(c *gin.Context) {
		for _, product := range *products {

			if strconv.Itoa(product.Id) == c.Param("id") {
				c.JSON(200, product)
				return
			}
		}

		c.String(404, "Can't find a product with the given id")
	}

}
