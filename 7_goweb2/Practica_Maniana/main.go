package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const TOKEN = "1234"

var products []Producto

type Producto struct {
	Id             int
	Nombre         string
	Color          string
	Precio         float64
	Stock          int
	Codigo         string
	Publicado      bool
	Fecha_creacion string
}

func getID(p []Producto) int {
	var clave int
	for _, v := range p {
		clave = v.Id
	}
	clave++
	return clave
}

func validCamp(p Producto) (camp string) {
	camp = ""
	switch {
	case p.Nombre == "":
		camp = "Nombre"
	case p.Color == "":
		camp = "Color"
	case p.Precio == 0:
		camp = "Precio"
	case p.Stock == 0:
		camp = "Stock"
	case p.Codigo == "":
		camp = "Codigo"
	case p.Fecha_creacion == "":
		camp = "Fecha_creacion"
	}
	return
}

func addProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var p Producto
		token := ctx.Request.Header.Get("token")
		if token != TOKEN {
			ctx.String(401, "no tiene permiso para realizar la peticion solicitada")
			return
		}

		if err := ctx.Bind(&p); err != nil {
			ctx.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}
		if camp := validCamp(p); camp != "" {
			ctx.String(400, fmt.Sprintf("El campo %v es requerido", camp))
			return
		}
		p.Id = getID(products)
		ctx.JSON(200, p)
		products = append(products, p)
	}
}

func main() {
	router := gin.Default()
	g := router.Group("/addproduct")
	{
		g.POST("/", addProduct())
	}
	router.Run()
}
