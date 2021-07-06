package main

import (
	"github.com/gin-gonic/gin"
	"github.com/josephcasa/meli_bootcamp10/handlers"
)

func RunServer() {
	router := gin.Default()

	router.GET("/greet", handlers.Greet)
	router.GET("/products", handlers.GetAll)
	router.GET("/filterProducts", handlers.FilterProducts)
	router.GET("/products/:id", handlers.GetProduct)

	router.Run()
}
