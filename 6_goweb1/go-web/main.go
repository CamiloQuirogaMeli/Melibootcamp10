package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

type producto struct {
	Id        int
	Nombre    string
	Color     string
	Precio    float64
	Stock     uint
	Codigo    string
	Publicado bool
	Fecha     string
}

func getAll(ctxt *gin.Context) {
	prods := []producto{}
	content, _ := ioutil.ReadFile("./products.json")

	_ = json.Unmarshal(content, &prods)

	var filtrados []producto

	status := ctxt.Query("color")
	fmt.Println("status", status)
	for _, prod := range prods {
		if status == string(prod.Color) {
			filtrados = append(filtrados, prod)
			fmt.Println("prod", prod)

		}
	}
	if len(filtrados) != 0 {
		ctxt.JSON(200, gin.H{
			"message": filtrados,
		})
		return
	}
	ctxt.JSON(404, gin.H{
		"message": prods,
	})
}

func getById(ctxt *gin.Context) {
	prods := []producto{}
	content, _ := ioutil.ReadFile("./products.json")

	_ = json.Unmarshal(content, &prods)

	var productoFiltrado producto
	userId, _ := strconv.Atoi(ctxt.Param("id"))

	for _, user := range prods {
		if user.Id == userId {
			productoFiltrado = user
		}
	}

	if productoFiltrado == (producto{}) {
		ctxt.String(404, "Información del empleado ¡No existe!")

		return
	}

	ctxt.JSON(200, gin.H{
		"producto": productoFiltrado,
	})
}

func main() {
	r := gin.Default()
	r.GET("/products", getAll)
	r.GET("/products/:id", getById)

	r.Run(":8081")
}
