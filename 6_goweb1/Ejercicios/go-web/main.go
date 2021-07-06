package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id              int
	Nombre          string
	Color           string
	Precio          float64
	Stock           int
	Codigo          string
	Publicado       bool
	FechaDeCreacion string
}

func ReadJson(c *gin.Context) []Producto {
	productos := []Producto{}

	data, err := ioutil.ReadFile("./products.json")

	if err != nil {
		log.Fatal("Error Data:", err)
		fmt.Println("Error Data:", err)
	}

	err = json.Unmarshal([]byte(data), &productos)

	// fmt.Println("Productos", productos)

	return productos
}

func GetAll(c *gin.Context) {
	productos := ReadJson(c)

	// c.String(200, "%v", productos)
	c.JSON(200, gin.H{
		"productos": productos,
	})
}

func FiltrarProductos(c *gin.Context) {
	productos := ReadJson(c)

	filtrados := []Producto{}

	for _, p := range productos {
		if c.Query("color") == p.Color {
			filtrados = append(filtrados, p)
		}
	}

	if len(filtrados) != 0 {
		c.JSON(200, gin.H{
			"filtrados": filtrados,
		})
		return
	}

	c.JSON(200, gin.H{
		"productos": productos,
	})
}

func FiltrarPorId(c *gin.Context) {
	productos := ReadJson(c)

	producto := Producto{}
	var indice int

	for i, p := range productos {
		if c.Param("id") == strconv.Itoa(p.Id) {
			indice = i
			producto = p
		}
	}

	if producto.Id != 0 {
		c.String(200, "Datos del producto con ID %v. Posicion del array %v - Nombre del Producto: %s", producto.Id, indice, producto.Nombre)
	} else {
		c.String(404, "El producto con ID: %s no existe.", c.Param("id"))
	}
}

func main() {

	router := gin.Default()

	// EjercicioTM -> Ejercicio 1
	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Mario",
		})
	})

	// EjercicioTM -> Ejercicio 2
	router.GET("/tematicas", GetAll)

	// EjerciciosTT -> Ejercicio 1
	router.GET("tematicas/filtro", FiltrarProductos)

	// EjerciciosTT -> Ejercicio 1
	router.GET("tematicas/:id", FiltrarPorId)

	log.Fatal(router.Run())
}
