package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/josephcasa/meli_bootcamp10/product"
)

func FilterProducts(products *[]product.Product) gin.HandlerFunc {

	return func(c *gin.Context) {
		filteredProducts := []product.Product{}

		for _, product := range *products {

			if FindMatchingProducts(c, &product) {
				filteredProducts = append(filteredProducts, product)
			}

		}
		c.JSON(200, filteredProducts)
	}
}

func FindMatchingProducts(c *gin.Context, product *product.Product) bool {

	match := (c.Query("id") == strconv.Itoa(product.Id) ||
		c.Query("name") == product.Name ||
		c.Query("color") == product.Color ||
		c.Query("code") == product.Code ||
		c.Query("creationDate") == product.CreationDate)

	return match
}
