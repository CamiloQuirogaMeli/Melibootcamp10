package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Products struct {
	Products []Product `json:"products"`
}
type Product struct {
	Id            string  `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_creacion"`
}

func main() {

	router := gin.Default()

	router.GET("/ping", func(g *gin.Context) {

		g.JSON(200, gin.H{
			"message": "Hola Luciano!",
		})
	})

	router.POST("/newproduct", nuevoProducto())

	router.Run()
}

var sProducts []Product
var incrementarId int = 0

func nuevoProducto() gin.HandlerFunc {
	return func(g *gin.Context) {
		var req Product
		token := g.Request.Header.Get("token")
		if token != "1234" {
			g.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}

		if err := g.Bind(&req); err != nil {
			g.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.Nombre == "" {
			g.String(400, "El campo Nombre es requerido")
			return
		} else if req.Color == "" {
			g.String(400, "El campo Color es requerido")
			return
		} else if req.Precio == 0 {
			g.String(400, "El campo Precio es requerido")
			return
		} else if req.Stock == 0 {
			g.String(400, "El campo Stock es requerido")
			return
		} else if req.Codigo == "" {
			g.String(400, "El campo Codigo es requerido")
			return
		} else if req.FechaCreacion == "" {
			g.String(400, "El campo Fecha es requerido")
			return
		}

		incrementarId++
		req.Id = strconv.Itoa(incrementarId)
		sProducts = append(sProducts, req)
		g.JSON(200, req)
		log.Println(sProducts)
	}
}
