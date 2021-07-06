package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id       int    `json:"id"`
	Cantidad int    `json:"cantidad" binding:"required"`
	Nombre   string `json:"nombre" binding:"required"`
	Precio   int    `json:"precio" binding:"required"`
}

var productos []Producto
var maxId int = 0

func main() {
	fmt.Println("Hola mundo")
	r := gin.Default()
	r.POST("/productos", crearProducto())
	r.Run(":8080")
}
func crearProducto() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Producto
		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		maxId++
		req.Id = maxId
		productos = append(productos, req)
		c.JSON(200, req)
	}
}
