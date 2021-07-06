package server

import (
	"github.com/gin-gonic/gin"
	"github.com/josephcasa/meli_bootcamp10/handlers"
	"github.com/josephcasa/meli_bootcamp10/product"
	"github.com/josephcasa/meli_bootcamp10/utilities"
)

var products []product.Product

const AUTH_TOKEN = "12345"

func RunServer() {

	var maxId int
	var err error

	if products, err = utilities.LoadFile("products.json"); err != nil {
		panic(err)
	}

	for _, p := range products {
		if maxId < p.Id {
			maxId = p.Id
		}
	}

	router := gin.Default()

	router.GET("/greet", handlers.Greet)

	productsGroup := router.Group("/products")

	productsGroup.GET("/", handlers.GetAll(&products))
	productsGroup.GET("/filter", handlers.FilterProducts(&products))
	productsGroup.GET("/:id", handlers.GetProduct(&products))
	productsGroup.POST("/", handlers.CreateProduct(AUTH_TOKEN, utilities.AddNewProduct(products, maxId)))

	router.Run()
}
