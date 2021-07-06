package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tomastropea/meli_bootcamp10/product"
	"strconv"
)

func FilterProductsHandler(ctx *gin.Context) {

	var products []product.Product
	var err error

	selectedProducts := make([]product.Product, 0)
	if products, err = LoadProductsFromFile(PRODUCTS_PATH); err != nil {

		ctx.String(400, err.Error())
		return
	} else {

		for _, product := range products {
			
			if QueryMatchesProduct(ctx, &product) {

				selectedProducts = append(selectedProducts, product)
			}
		}

		ctx.JSON(200, selectedProducts)
	}
}

func QueryMatchesProduct(ctx *gin.Context, product *product.Product) bool {
	return 	ctx.Query("id") 			== strconv.Itoa(product.ID) 				|| 
			ctx.Query("name") 			== product.Name 							||
			ctx.Query("color") 			== product.Color 							||
			ctx.Query("price") 			== strconv.Itoa(product.Price) 				|| 
			ctx.Query("stock") 			== strconv.Itoa(product.Stock) 				||
			ctx.Query("code")  			== product.Code 							||
			ctx.Query("published") 		== strconv.FormatBool(product.Published) 	|| 
			ctx.Query("creationDate")	== product.CreationDate 
}