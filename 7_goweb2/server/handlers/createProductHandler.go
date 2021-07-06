package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tomastropea/meli_bootcamp10/product"
	"fmt"
)

func CreateProductHandler(AUTH_TOKEN string, addNewProductToDatabase func(product.Product) (product.Product, error)) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var product product.Product
		var err		error

		authToken := ctx.Request.Header.Get("token")
		if authToken != AUTH_TOKEN {
			ctx.JSON(401, gin.H{
				"error" : "no tiene permisos para realizar la petición solicitada",
			})
			return
		}


		if err := ctx.Bind(&product); err != nil {
			ctx.String(400, "The product data is invalid")
			return
		}
		if field := findInvalidFieldForProduct(product); field != "" {
			ctx.String(400, fmt.Sprintf("El campo %s es requerido", field))
			return
		}

		product, err = addNewProductToDatabase(product)

		if err != nil {
			ctx.String(500, "Internal server error when saving product to database")
		} else {
			ctx.JSON(200, product)
		}
	}
}

func findInvalidFieldForProduct(p product.Product) string {
	
	invalidField := ""

	switch  {
		case p.Name == "": invalidField = "name"
		case p.Color == "": invalidField = "color"
		case p.Price == 0: invalidField = "price"
		case p.Stock == 0: invalidField = "stock"
		case p.Code == "": invalidField = "code"
		case p.CreationDate == "": invalidField = "creationDate"
	}

	return invalidField
}