package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type producto struct {
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
		//creamos nuestro server y router
		router := gin.Default()

		//creamos nuestro Handler
		router.GET("/hola", func (c *gin.Context){//http://localhost:8080/hola

			c.JSON(200, gin.H{
				"message": "Hola Ignacio",
			})
		})

		router.GET("/productos", func (c *gin.Context){//http://localhost:8080/productos

			dat, err := ioutil.ReadFile("./6_goweb1/TM/Ejercicio1/products.json")

			if err != nil{
				c.JSON(400, gin.H{
					"error": "problemas al leer el archivo solicitado.",
				})
			}else{

			var prods []producto

			json.Unmarshal(dat, &prods)
			fmt.Println(prods)
			c.String(200, "%v", prods)

			}
		})
		router.Run()
}
