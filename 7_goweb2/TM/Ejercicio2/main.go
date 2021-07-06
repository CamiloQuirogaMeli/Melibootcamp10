package main

import (
	"fmt"
	"strings"

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
	fmt.Println("App Giuliano")
	r := gin.Default()
	r.POST("/productos", crearProducto())
	r.Run(":8080")
}
func crearProducto() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Producto
		if err := c.Bind(&req); err != nil {
			err1 := strings.Split(err.Error(), "'")
			fmt.Print(err1)
			c.String(400, "el campo %s es requerido", err1[1])
			return
		}
		maxId++
		req.Id = maxId
		productos = append(productos, req)
		c.JSON(200, req)
	}
}
