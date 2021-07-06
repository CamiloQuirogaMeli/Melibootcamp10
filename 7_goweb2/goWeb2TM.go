package main

import (
	"github.com/gin-gonic/gin"
)

//GO WEB 2

type Request struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json: "tipo"`
	Cantidad int     `json: "cantidad"`
	Precio   float64 `json: "precio"`
}

var productos []Request
var maxID int

func main() {
	//CON GRUPOS
	r := gin.Default()
	pr := r.Group("/productos")
	pr.POST("/", Guardar())

	/* SIN GRUPOS
	r.POST("/productos", func(c *gin.Context) {
		var req Request
		if err := c.Bind(&req); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}
		req.ID = 4
		c.JSON(200, req)
	})
	*/ //SIN GRUPOS
	r.Run()
}

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}
		var req Request
		if err := c.Bind(&req); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}
		maxID++
		req.ID = maxID
		productos = append(productos, req)
		c.JSON(200, req)
	}
}
