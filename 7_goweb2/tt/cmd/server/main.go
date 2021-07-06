package main

import (
	"github.com/conavarr/tt/cmd/server/handler"
	"github.com/conavarr/tt/internal/products"
	"github.com/gin-gonic/gin"
)

func main(){
	repository := products.NewRepository()
	service := products.NewService(repository)
	p := handler.NewProductController(service)


	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("", p.Store())
	pr.GET("", p.GetAll())
	r.Run()
}