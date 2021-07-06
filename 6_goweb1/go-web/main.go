package main

import (
	"github.com/gin-gonic/gin"
)

type handler struct {
	getOne func(context *gin.Context)
	getAll func(context *gin.Context)
}

func getAll(context *gin.Context) {

	var products []producto = getProducts()

	var isActive string = context.Query("active")

	if isActive == "" {

		context.JSON(200, gin.H{"message": products})

	} else {

		queryFilter(isActive, &products, context)
	}

}

func getOne(context *gin.Context) {

	var products []producto = getProducts()

	filterOne(context, &products)

}

func main() {

	handlers := handler{
		getOne: getOne,
		getAll: getAll,
	}

	r := gin.Default()

	r.GET("/tematicas", handlers.getAll)
	r.GET("/tematicas/:id", handlers.getOne)

	r.Run()
}
