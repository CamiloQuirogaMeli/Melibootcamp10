package main

import (
	"encoding/json"
	"fmt"

	//"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	//"log"
)

type Producto struct {
	Id int
	Nombre string
	Color string
	Precio float64
	Stock uint
	Codigo string
	Publicado bool
	FechaCreacion string
}

func main(){

	/* SABER DIR DONDE ESTOY PARADO
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	*/

	// Creo servidor gin
	router := gin.Default()

	// Creo Handler EJ2
	router.GET("/hola", func(c *gin.Context){ //http://localhost:8080/hola

		c.JSON(200, gin.H{
			"message":"Hola Ignacio!",
		})
	})

	//EJ 3
	router.GET("/productos", func(c *gin.Context) { //http://localhost:8080/productos
		dataJson, err := ioutil.ReadFile("./6_goweb1/TM/Ejercicio1/products.json")

		if err != nil {
			c.JSON(400, gin.H{
				"error": "problemas al leer el archivo solicitado.",
			})
		}else{
			var productos []Producto

			_ = json.Unmarshal(dataJson, &productos)
			fmt.Println(productos)
			c.String(200, "%v", productos)
			//fmt.Println(productos)
			}

	})

	router.Run()
// https://github.com/extmatperez/meli_bootcamp10/blob/demian_gonzalo/6_goweb1/go-web/main.go
}
