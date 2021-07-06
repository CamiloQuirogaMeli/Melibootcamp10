package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type request struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float32 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_creacion"`
}

var products []request
var req request

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}

		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := validarCampos(); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			products = append(products, req)
			indice := len(products)
			req.Id = products[indice-1].Id + 1

			c.JSON(200, req)
		}

	}
}

func validarCampos() error {
	if req.Codigo == "" || req.Color == "" {
		return errors.New("falla validacion")
	}
	return nil
}

func main() {
	router := gin.Default()
	router.POST("/productos", Guardar())
	router.Run()
}
