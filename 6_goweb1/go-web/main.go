package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/hello-world", helloWorld)

	productsRouter := router.Group("/products")
	{
		productsRouter.GET("/", getProducs)
		productsRouter.GET("/:id", getOneProduct)
	}
	router.Run()
}

func helloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hola Willian!",
	})
}

func getProducs(c *gin.Context) {
	products, status := getAllProducts()
	var filtrados, prevFilter []Product

	if c.Query("id") != "" {
		queryValue, _ := strconv.Atoi(c.Query("id"))
		for _, p := range products {
			if queryValue == p.Id {
				prevFilter = append(prevFilter, p)
			}
		}
	} else {
		prevFilter = products
	}

	if c.Query("nombre") != "" {
		for _, p := range prevFilter {
			if strings.Contains(p.Nombre, c.Query("nombre")) {
				filtrados = append(filtrados, p)
			}
		}
		prevFilter = filtrados
		filtrados = filtrados[:0]
	}

	if c.Query("color") != "" {
		for _, p := range prevFilter {
			if c.Query("color") == p.Color {
				filtrados = append(filtrados, p)
			}
		}
		prevFilter = filtrados
		filtrados = filtrados[:0]
	}

	if c.Query("precio") != "" {
		queryValue, _ := strconv.ParseFloat(c.Query("precio"), 64)

		for _, p := range prevFilter {
			if queryValue == p.Precio {
				filtrados = append(filtrados, p)
			}
		}
		prevFilter = filtrados
		filtrados = filtrados[:0]
	}

	if c.Query("stock") != "" {
		queryValue, _ := strconv.Atoi(c.Query("stock"))

		for _, p := range prevFilter {
			if queryValue == p.Stock {
				filtrados = append(filtrados, p)
			}
		}
		prevFilter = filtrados
		filtrados = filtrados[:0]
	}

	if c.Query("codigo") != "" {
		for _, p := range prevFilter {
			if c.Query("codigo") == p.Codigo {
				filtrados = append(filtrados, p)
			}
		}
		prevFilter = filtrados
		filtrados = filtrados[:0]
	}

	if c.Query("publicado") != "" {
		queryValue, _ := strconv.ParseBool(c.Query("publicado"))
		for _, p := range prevFilter {
			if queryValue == p.Publicado {
				filtrados = append(filtrados, p)
			}
		}
		prevFilter = filtrados
		filtrados = filtrados[:0]
	}

	if c.Query("fecha_creacion") != "" {
		for _, p := range prevFilter {
			if c.Query("fecha_creacion") == p.FechaCreacion {
				filtrados = append(filtrados, p)
			}
		}
		prevFilter = filtrados
		filtrados = filtrados[:0]
	}

	if status != 200 {
		c.JSON(status, gin.H{
			"error": "ops ha ocurrido un error",
		})
	} else {
		if prevFilter != nil {
			c.JSON(200, gin.H{
				"data": prevFilter,
			})
		} else {
			c.JSON(200, gin.H{
				"data": []Product{},
			})
		}
	}
}

func getAllProducts() ([]Product, int) {
	dat, err := ioutil.ReadFile("./products.json")
	if err != nil {
		return []Product{}, 500
	}

	var myProducts []Product

	json.Unmarshal(dat, &myProducts)
	return myProducts, 200
}

func getOneProduct(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	dat, err := ioutil.ReadFile("./products.json")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "ops ha ocurrido un problema",
		})
		return
	}

	var myProducts []Product

	json.Unmarshal(dat, &myProducts)
	for _, product := range myProducts {
		if product.Id == id {
			c.JSON(200, gin.H{
				"product": product,
			})
			return
		}
	}
	c.JSON(500, gin.H{
		"error": "producto no encontrado",
	})
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
