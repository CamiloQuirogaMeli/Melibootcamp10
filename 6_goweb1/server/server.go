package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tomastropea/meli_bootcamp10/server/handlers"
)

func RunServer() {
	router  := gin.Default()

	router.GET("/greetUser", handlers.GreetUserHandler)
	router.GET("/products", handlers.GetAllHandler)
	router.GET("/filterProducts", handlers.FilterProductsHandler)
	router.GET("/products/:id", handlers.GetProductByIdHandler)

	router.Run()
}
 