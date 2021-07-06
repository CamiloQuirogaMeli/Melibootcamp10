package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creamos el router
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Lucas!",
		})
	})

	router.GET("/products/getAll", func(c *gin.Context) {
		prod, err := ioutil.ReadFile("products.json")

		if err == nil {
			keys := make([]Product, 0)
			json.Unmarshal(prod, &keys)
			c.JSON(200, gin.H{
				"message": keys,
			})
		} else {
			c.JSON(400, gin.H{
				"message": "Error al leer el archivo",
			})

		}
	})

	router.GET("/products/:id", func(c *gin.Context) {
		prod, err := ioutil.ReadFile("products.json")
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Error al leer el archivo",
			})
		}
		keys := make([]Product, 0)
		var prods []Product
		json.Unmarshal(prod, &keys)

		if c.Param("id") == "" {
			// FilterId(c.Param("id"), keys, &prods)
		}

		if len(prods) > 0 {
			c.JSON(200, gin.H{
				"message": prods,
			})
		} else {
			c.JSON(404, gin.H{
				"message": c.Param("id"),
			})
		}

	})

	router.Run(":8081")
}

type Product struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_creacion"`
}

func FilterId(id int, origin []Product, destination *[]Product) {

	for _, product := range origin {
		if product.Id == id {
			*destination = append(*destination, product)
		}
	}
}
