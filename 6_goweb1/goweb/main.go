package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id              int
	Nombre          string `from:"nombre" json:"nombre"`
	Color           string
	Precio          int
	Stock           int
	Codigo          string
	Publicado       bool
	FechaDeCreacion string
}

func leerArchivo(c *gin.Context) []Producto {
	productos := []Producto{}

	data, err := ioutil.ReadFile("./productos.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	if err := json.Unmarshal([]byte(data), &productos); err != nil {
		log.Fatal(err)
	}
	return productos
}

func GetAll(c *gin.Context) {
	productos := leerArchivo(c)
	c.String(200, "%v", productos)

}

func GetFilterByName(c *gin.Context) {
	productosFiltrados := filtrarProductos(c)
	c.String(200, "%v", productosFiltrados)

}

func filtrarProductos(c *gin.Context) []Producto {
	productos := leerArchivo(c)
	productosFiltrados := []Producto{}
	var encontroProducto bool

	for _, e := range productos {
		if c.Query("nombre") == e.Nombre {
			productosFiltrados = append(productosFiltrados, e)
			encontroProducto = true
		}
	}
	if !encontroProducto {
		log.Fatal("No se encontraron productos")
	}
	return productosFiltrados

}

func Saludo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hola gaston",
	})
}

func main() {

	router := gin.Default()

	router.GET("/hola", Saludo)

	router.GET("/productos", GetAll)
	router.GET("/productosFiltrados", GetFilterByName)
	router.Run()
}
