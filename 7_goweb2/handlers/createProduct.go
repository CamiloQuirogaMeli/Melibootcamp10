package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/josephcasa/meli_bootcamp10/product"
)

func CreateProduct(AUTH_TOKEN string, addNewProduct func(product.Product) (product.Product, error)) gin.HandlerFunc {

	return func(c *gin.Context) {

		var product product.Product
		var err error

		authToken := c.Request.Header.Get("token")

		if authToken != AUTH_TOKEN {
			c.JSON(401, gin.H{
				"error": "You do not have permission to make the request",
			})
			return
		}

		if err := c.Bind(&product); err != nil {
			c.String(404, "The product data is invalid")
			return
		}
		if field := FindInvalidField(product); field != " " {
			c.String(404, fmt.Sprintf("You must provide valid input for %s", field))
			return
		}

		product, err = addNewProduct(product)

		if err != nil {
			c.String(500, "Internal server error")
		} else {
			c.JSON(200, product)
		}
	}
}

func FindInvalidField(product product.Product) string {

	var invalidField string

	switch {
	case product.Name == "":
		invalidField = "name"
	case product.Color == "":
		invalidField = "color"
	case product.Price == 0:
		invalidField = "price"
	case product.Stock == 0:
		invalidField = "stock"
	case product.Code == "":
		invalidField = "code"
	case product.CreationDate == "":
		invalidField = "creationDate"
	default:
		invalidField = " "
	}

	return invalidField
}
