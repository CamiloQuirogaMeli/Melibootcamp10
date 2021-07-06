package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

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

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Nina!",
		})
	})

	products := router.Group("/productos")
	{
		products.GET("/", GetAll)
		products.GET("/:id", GetById)
	}

	router.Run()
}

func GetAll(c *gin.Context) {
	dataProductos, err1 := ioutil.ReadFile("./products.json")
	if err1 != nil {
		log.Fatal(err1)
	} else {
		var productos []Producto
		err2 := json.Unmarshal(dataProductos, &productos)
		if err2 != nil {
			log.Fatal(err2)
		} else {
			productos = filtrarProductos(productos, c)
			c.JSON(200, gin.H{
				"cantidad":  len(productos),
				"productos": productos,
			})
		}
	}
}

func GetById(c *gin.Context) {
	dataProductos, err1 := ioutil.ReadFile("./products.json")
	if err1 != nil {
		log.Fatal(err1)
	} else {
		var productos []Producto
		err2 := json.Unmarshal(dataProductos, &productos)
		if err2 != nil {
			log.Fatal(err2)
		} else {
			var producto = Producto{Id: -1}
			for _, p := range productos {
				id, _ := strconv.Atoi(c.Param("id"))
				if p.Id == id {
					producto = p
				}
			}
			if producto.Id != -1 {
				c.JSON(200, gin.H{
					"detalle": producto,
				})
			} else {
				c.String(404, "Error, no se encontró el producto")
			}
		}
	}
}

func filtrarProductos(productos []Producto, c *gin.Context) []Producto {
	var productosFiltrados []Producto
	for _, producto := range productos {
		pasaElFiltro := true
		if len(c.Query("nombre")) > 0 && c.Query("nombre") != producto.Nombre { // filtro nombre
			pasaElFiltro = false
		}
		if len(c.Query("color")) > 0 && c.Query("color") != producto.Color { // filtro color
			pasaElFiltro = false
		}
		if len(c.Query("precio")) > 0 { // filtro precio
			precio, _ := strconv.ParseFloat(c.Query("precio"), 64)
			if precio != producto.Precio {
				pasaElFiltro = false
			}
		}
		if len(c.Query("stock")) > 0 { // filtro stock
			stock, _ := strconv.Atoi(c.Query("stock"))
			if stock != producto.Stock {
				pasaElFiltro = false
			}
		}
		if len(c.Query("codigo")) > 0 && c.Query("codigo") != producto.Codigo { // filtro codigo
			pasaElFiltro = false
		}
		if len(c.Query("publicado")) > 0 { // filtro publicado
			publicado, _ := strconv.ParseBool(c.Query("publicado"))
			if publicado != producto.Publicado {
				pasaElFiltro = false
			}
		}
		if len(c.Query("fechaCreacion")) > 0 && c.Query("fechaCreacion") != producto.FechaCreacion { // filtro fecha de creacion
			pasaElFiltro = false
		}
		if pasaElFiltro {
			productosFiltrados = append(productosFiltrados, producto)
		}
	}
	return productosFiltrados
}
