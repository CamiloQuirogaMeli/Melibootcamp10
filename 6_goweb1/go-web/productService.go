package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func getProducts() (products []producto) {

	dataFromFile, errorFromReadingFile := ioutil.ReadFile("./products.json")

	if errorFromReadingFile != nil {

		log.Fatal(errorFromReadingFile)

	} else {

		var errorFromUnmarshal error = json.Unmarshal(dataFromFile, &products)

		if errorFromUnmarshal != nil {
			log.Fatal(errorFromUnmarshal)
		}

	}

	return

}

func queryFilter(isActive string, products *[]producto, context *gin.Context) {

	productos := *products

	active, err := strconv.ParseBool(isActive)

	var productsFiltrados []producto = []producto{}

	if err == nil {
		for _, value := range productos {
			if value.Publicado == active {
				productsFiltrados = append(productsFiltrados, value)
			}

		}

		context.JSON(200, gin.H{"message": productsFiltrados})

	} else {
		log.Fatal(err)
	}
}

func filterOne(context *gin.Context, products *[]producto) {

	productos := *products

	param, error := strconv.Atoi(context.Param("id"))

	if error == nil {
		for _, value := range productos {
			if param == value.Id {

				context.JSON(200, gin.H{"message": value})
			}
		}
	}
}
