package main

import (
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

func filtrarProductos(ctx *gin.Context, productos []Producto) []Producto {
	var filtrados []Producto
	for _, producto := range productos {
		if ctx.Query("nombre") == producto.Nombre {
			filtrados = append(filtrados, producto)
		}
		if ctx.Query("color") == producto.Color {
			filtrados = append(filtrados, producto)
		}
		if ctx.Query("precio") == strconv.FormatFloat(producto.Precio, 'f', -1, 64) {
			filtrados = append(filtrados, producto)
		}
		if ctx.Query("stock") == strconv.Itoa(producto.Stock) {
			filtrados = append(filtrados, producto)
		}
		if ctx.Query("codigo") == producto.Codigo {
			filtrados = append(filtrados, producto)
		}
		if ctx.Query("publicado") == strconv.FormatBool(producto.Publicado) {
			filtrados = append(filtrados, producto)
		}
		if ctx.Query("fecha") == producto.FechaCreacion {
			filtrados = append(filtrados, producto)
		}
	}
	return filtrados
}

func listaProductos() []Producto {
	p1 := Producto{1, "Remera", "Rojo", 15.90, 10, "ABC123", false, "20/05/2020"}
	p2 := Producto{2, "Pantalones", "Negro", 150.00, 1, "DEF456", true, "25/05/2021"}
	productos := []Producto{p1, p2}
	return productos
}

func getByID(ctx *gin.Context) {
	productos := listaProductos()
	var selected Producto
	for _, producto := range productos {
		if ctx.Param("id") == strconv.Itoa(producto.Id) {
			selected = producto
			break
		}
	}
	if selected.Id != 0 {
		ctx.JSON(200, gin.H{"Producto": selected})
	} else {
		ctx.JSON(404, gin.H{"Mensaje:": "No hay datos disponibles"})
	}
}

func obtenerProductos(ctx *gin.Context) {
	productos := listaProductos()
	if len(ctx.Request.URL.Query()) > 0 {
		productos = filtrarProductos(ctx, productos)
	}
	if len(productos) > 0 {
		ctx.JSON(200, gin.H{"Productos": productos})
	} else {
		ctx.JSON(404, gin.H{"Mensaje:": "No hay datos disponibles"})
	}
}

func main() {
	router := gin.Default()

	router.GET("/hello-ariel", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola Ariel"})
	})

	productos := router.Group("/productos")
	{
		productos.GET("/", obtenerProductos)
		productos.GET("/:id", getByID)
	}

	router.Run()
}
