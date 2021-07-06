package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id        int
	Nombre    string
	Precio    int
	Publicado bool
	Color     string
	Fecha     string
	Stock     int
	Codigo    string
}

func GetAll(direccion string, p *[]product) []product {
	content, err := ioutil.ReadFile(direccion)
	if err != nil {
		log.Fatal(err)
	}

	err1 := json.Unmarshal(content, p)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Print(*p)
	return *p
}

func main() {

	productos := []product{}

	router := gin.Default()
	router.GET("/productos", func(c *gin.Context) {

		c.String(http.StatusOK, "%v", GetAll("../Ejercicio1/productos.json", &productos))

	})
	router.Run()
}
