package main

import (
	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre" binding:"required"`
	Color         string  `json:"color" binding:"required"`
	Precio        float64 `json:"precio" binding:"required,numeric,min=1"`
	Stock         int     `json:"stock" binding:"required"`
	Codigo        string  `json:"codigo" binding:"required"`
	Publicado     bool    `json:"publicado" binding:"required"`
	FechaCreacion string  `json:"fecha_creacion" binding:"required"`
}

var productos []Producto

func createProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la petición solicitada",
			})
			return
		}

		var req Producto

		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		req.Id = len(productos) + 1
		productos = append(productos, req)

		ctx.JSON(200, req)
	}
}

func main() {
	r := gin.Default()

	r.POST("/productos", createProduct())

	r.Run()
}
