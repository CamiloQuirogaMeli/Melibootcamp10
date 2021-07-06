package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Products struct {
	Products []Product `json:"products"`
}
type Product struct {
	Id            string  `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_creacion"`
}

func main() {

	router := gin.Default()

	router.GET("/ping", func(g *gin.Context) {

		g.JSON(200, gin.H{
			"message": "Hola Luciano!",
		})
	})
	router.GET("/products", GetAll)
	router.GET("/products/id/:id", filtrarProductoId)
	router.GET("/products/name/:nombre", filtrarProductoNombre)
	router.GET("/products/color/:color", filtrarProductoColor)
	router.GET("/products/filtrar", buscarProducto)

	router.Run()
}

func GetAll(g *gin.Context) {
	jsonFile, err := os.Open("./products.js")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("archivo abierto")
	defer jsonFile.Close()

	// leo el archivo abierto
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//inicializo el array de productos
	var products Products

	//obtenemos todos los productos del byteArray
	json.Unmarshal(byteValue, &products)

	g.String(200, "%v", products)

}

func filtrarProductoId(g *gin.Context) {
	jsonFile, err := os.Open("./products.js")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("archivo abierto")
	defer jsonFile.Close()

	// leo el archivo abierto
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//inicializo el array de productos
	var products Products

	//obtenemos todos los productos del byteArray
	json.Unmarshal(byteValue, &products)

	existe := false
	var producto Product
	for _, v := range products.Products {
		if v.Id == g.Param("id") {
			producto, existe = v, true
		}
	}

	if existe {
		g.String(200, "%v", producto)
	} else {
		messageNotFound := "No existe el producto seleccionado!"
		g.String(404, "%v", messageNotFound)
	}
}

func filtrarProductoNombre(g *gin.Context) {
	jsonFile, err := os.Open("./products.js")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("archivo abierto")
	defer jsonFile.Close()

	// leo el archivo abierto
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//inicializo el array de productos
	var products Products

	//obtenemos todos los productos del byteArray
	json.Unmarshal(byteValue, &products)

	existe := false
	var producto Product
	for _, v := range products.Products {
		if v.Nombre == g.Param("nombre") {
			producto, existe = v, true
		}
	}

	if existe {
		g.String(200, "%v", producto)
	} else {
		messageNotFound := "No existe el producto seleccionado!"
		g.String(404, "%v", messageNotFound)
	}
}

func filtrarProductoColor(g *gin.Context) {
	jsonFile, err := os.Open("./products.js")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("archivo abierto")
	defer jsonFile.Close()

	// leo el archivo abierto
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//inicializo el array de productos
	var products Products

	//obtenemos todos los productos del byteArray
	json.Unmarshal(byteValue, &products)

	existe := false
	var producto Product
	for _, v := range products.Products {
		if v.Color == g.Param("color") {
			producto, existe = v, true
		}
	}

	if existe {
		g.String(200, "%v", producto)
	} else {
		messageNotFound := "No existe el producto seleccionado!"
		g.String(404, "%v", messageNotFound)
	}
}

func buscarProducto(g *gin.Context) {
	jsonFile, err := os.Open("./products.js")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("archivo abierto")
	defer jsonFile.Close()

	// leo el archivo abierto
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//inicializo el array de productos
	var products Products

	//obtenemos todos los productos del byteArray
	json.Unmarshal(byteValue, &products)

	var filtrados []Product
	cantidadFiltrados := 0
	for _, v := range products.Products {
		if g.Query("nombre") == v.Nombre {
			filtrados = append(filtrados, v)
			cantidadFiltrados++
		}
	}

	if cantidadFiltrados > 0 {
		g.String(200, "Estos son los resultados filtrados! \n %v", filtrados)
		for _, v := range filtrados {
			log.Println("========Filtros por Query =========")
			log.Println(v.Id)
			log.Println(v.Nombre)
			log.Println(v.Color)
			log.Println(v.Precio)
		}
	} else {
		g.String(404, "Nada para filtrar!")
	}

}
