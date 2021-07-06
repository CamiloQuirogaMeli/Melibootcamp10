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
	fmt.Println("Hola mundo")
	r := gin.Default()
	r.POST("/productos", crearProducto())
	r.Run(":8080")
}
func crearProducto() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Producto
		token := c.Request.Header.Get("token")
		if token != "123456" {
			if token == "" {
				c.JSON(401, gin.H{
					"error": "no envió token",
				})
			} else {
				c.JSON(401, gin.H{
					"error": "token inválido",
				})
			}
		} else {

			if err := c.Bind(&req); err != nil {
				err1 := strings.Split(err.Error(), "'")
				c.String(400, "el campo %s es requerido", err1[1])
				return
			}
			maxId++
			req.Id = maxId
			productos = append(productos, req)
			c.JSON(200, req)
		}

	}

}
