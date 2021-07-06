package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre" binding:"required"`
	Color         string  `json:"color" binding:"required"`
	Precio        float64 `json:"precio" binding:"required"`
	Stock         int     `json:"stock" binding:"required"`
	Codigo        string  `json:"codigo" binding:"required"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_creacion" binding:"required"`
}

var products []product

func getNextID() int {
	if len(products) == 0 {
		return 1
	}
	return products[len(products)-1].Id + 1
}

func main() {

	router := gin.Default()

	productsRouter := router.Group("/products")
	{
		productsRouter.POST("/", addProduct())
	}

	router.Run()
}

func addProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}
		var p product
		if err := c.Bind(&p); err != nil {
			fmt.Println(p)
			if p.Nombre == "" {
				c.JSON(400, gin.H{
					"error": "el nombre del producto es requerido",
				})
				return
			}
			if p.Color == "" {
				c.JSON(400, gin.H{
					"error": "el color del producto es requerido",
				})
				return
			}

			if p.Precio == 0 {
				c.JSON(400, gin.H{
					"error": "el precio del producto es requerido",
				})
				return
			}
			if p.Stock == 0 {
				c.JSON(400, gin.H{
					"error": "el stock del producto es requerido",
				})
				return
			}
			if p.Codigo == "" {
				c.JSON(400, gin.H{
					"error": "el codigo del producto es requerido",
				})
				return
			}
			if p.FechaCreacion == "" {
				c.JSON(400, gin.H{
					"error": "la fecha de creacion del producto es requerida",
				})
				return
			}
		}
		if p.Precio <= 0 {
			c.JSON(400, gin.H{
				"error": "el precio del producto no puede ser negativo",
			})
			return
		}
		if p.Stock <= 0 {
			c.JSON(400, gin.H{
				"error": "el stock del producto no puede ser negativo",
			})
			return
		}
		p.Id = getNextID()
		products = append(products, p)
		c.JSON(200, p)

	}
}
