package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}

var products []Producto

func main() {
	dataProductos, err := ioutil.ReadFile("./products.json")
	if err != nil {
		log.Fatal(err)
	} else {
		json.Unmarshal(dataProductos, &products)
	}

	router := gin.Default()

	router.POST("/productos", Create())

	router.Run()
}

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Producto
		var errors []string
		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		if c.Request.Header.Get("token") != "123123" {
			c.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la petición solicitada",
			})
			return
		}
		if req.Nombre == "" {
			errors = append(errors, "el campo nombre es requerido")
		}
		if req.Color == "" {
			errors = append(errors, "el campo color es requerido")
		}
		if req.Precio == 0 {
			errors = append(errors, "el campo precio es requerido")
		}
		if req.Stock == 0 {
			errors = append(errors, "el campo stock es requerido")
		}
		if req.Codigo == "" {
			errors = append(errors, "el campo código es requerido")
		}
		if req.FechaCreacion == "" {
			errors = append(errors, "el campo fecha de creación es requerido")
		}
		if len(errors) > 0 {
			c.JSON(400, gin.H{
				"errors": errors,
			})
			return
		}

		if len(products) > 0 {
			req.Id = products[len(products)-1].Id + 1
		} else {
			req.Id = 1
		}

		products = append(products, req)
		c.JSON(200, req)
	}
}
