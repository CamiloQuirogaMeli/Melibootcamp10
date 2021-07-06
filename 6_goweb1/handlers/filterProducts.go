package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/josephcasa/meli_bootcamp10/products"
)

func FilterProducts(c *gin.Context) {

	var err error
	filteredProducts := []products.Product{}
	products := []products.Product{}

	if products, err = LoadFile("products.json"); err != nil {

		c.String(404, err.Error())
		return

	} else {

		for _, product := range products {

			if FindMatchingProducts(c, &product) {
				filteredProducts = append(filteredProducts, product)
			}

		}
		c.JSON(200, filteredProducts)
	}
}

func FindMatchingProducts(c *gin.Context, product *products.Product) bool {

	match := (c.Query("id") == strconv.Itoa(product.Id) ||
		c.Query("name") == product.Name ||
		c.Query("color") == product.Color ||
		c.Query("code") == product.Code ||
		c.Query("creationDate") == product.CreationDate)

	return match
}
