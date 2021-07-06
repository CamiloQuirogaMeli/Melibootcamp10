package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {

	products, err := LoadFile("products.json")

	if err != nil {
		c.String(500, err.Error())
		return
	}

	for _, product := range products {

		if strconv.Itoa(product.Id) == c.Param("id") {
			c.JSON(200, product)
			return
		}
	}

	c.String(404, "Can't find a product with the given id")
}
