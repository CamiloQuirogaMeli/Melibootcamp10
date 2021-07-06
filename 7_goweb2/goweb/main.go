package main

import (
	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id              int    `from:"id" json:"id"`
	Nombre          string `from:"nombre" json:"nombre" binding:"required"`
	Color           string `from:"color" json:"color" binding:"required"`
	Precio          int    `from:"precio" json:"precio" binding:"required"`
	Stock           int    `from:"stock" json:"stock" binding:"required"`
	Codigo          string `from:"codigo" json:"codigo" binding:"required"`
	Publicado       bool   `from:"publicado" json:"publicado" binding:"required"`
	FechaDeCreacion string `from:"fecha" json:"fecha" binding:"required"`
}

var productos []Producto

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "556677" {
			c.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}
		var req Producto
		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		req.Id = generarId()
		productos = append(productos, req)
		c.JSON(200, req)
	}
}

func generarId() int {
	var id int
	if len(productos) > 0 {
		id = productos[len(productos)-1].Id
	} else {
		id = 0
	}
	id++
	return id
}

func main() {
	router := gin.Default()
	prod := router.Group("/productos")
	prod.POST("/", Guardar())
	router.Run()
}
