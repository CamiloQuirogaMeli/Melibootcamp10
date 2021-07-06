package main

import (
	"github.com/gin-gonic/gin"
)

var products = []Product{}

func main() {

	server := gin.Default()

	// server.POST("/products", func(ctx *gin.Context) {
	// 	var product Product

	// 	if err := ctx.Bind(&product); err != nil {
	// 		ctx.JSON(400, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	product.Id = 1

	// 	ctx.JSON(200, product)
	// })

	// //Another kind of handler
	apiProducts := server.Group("/products")
	apiProducts.GET("/", GetAll())
	apiProducts.POST("/", Save())

	server.Run()
}

func GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, products)
	}
}

func Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if token := ctx.Request.Header.Get("token"); token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "you haven't permissions for make the request",
			})
			return
		}

		var product Product

		if err := ctx.Bind(&product); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		errors := entityValidation(product)
		if len(errors) > 0 {
			ctx.JSON(400, errors)
			return
		}

		product.Id = findMaxId(products)
		products = append(products, product)

		ctx.JSON(200, product)
	}
}

func findMaxId(products []Product) int {
	if len(products) == 0 {
		return 1
	}
	p := products[len(products)-1]
	return p.Id + 1
}

func entityValidation(product Product) map[string]string {
	typeProductAccepteds := []string{"Electronic", "Smartphone", "Tools", "Gamer"}
	errors := make(map[string]string)

	if product.Id != 0 {
		errors["id"] = "its not required"
	}

	if product.Name == "" {
		errors["name"] = "cannot be empty"
	}

	if product.Price <= 0.0 {
		errors["price"] = "cannot be zero or negative"
	}

	if product.Quantity <= 0 {
		errors["quantity"] = "cannot be zero or negative"
	}

	if product.TypeProduct != "" {
		productAccepted := false
		for _, typeProduct := range typeProductAccepteds {
			if product.TypeProduct == typeProduct {
				productAccepted = true
				break
			}
		}

		if !productAccepted {
			errors["type"] = "not accepted. Type must be \"Electronic\", \"Smartphone\", \"Tools\" or \"Gamer\""
		}
	}

	return errors
}

// func Autentication() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		if token := ctx.Request.Header.Get("token"); token != "123456" {
// 			ctx.JSON(401, gin.H{
// 				"error": "Token not found",
// 			})
// 			return
// 		}
// 	}
// }

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	TypeProduct string  `json:"type"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
