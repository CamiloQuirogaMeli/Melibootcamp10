package main

import (
	"encoding/json"

	products "github.com/extmatperez/meli_bootcamp10/tree/martinez_federico/7_goweb2/TT/internal/product"
	"github.com/gin-gonic/gin"
)

var personas []products.Producto = []products.Producto{}

func handlerPersona() gin.HandlerFunc {

	return func(context *gin.Context) {

		token := context.Request.Header.Get("token")

		if token != "12345" {
			context.JSON(404, gin.H{
				"error": "No tiene permisos para realizar la peticion",
			})
			return
		}

		var request products.Producto

		var maxId int = 0 //repo

		if bindingError := context.Bind(&request); bindingError != nil {
			context.JSON(404, gin.H{
				"error": bindingError.Error(),
			})
			return
		}

		var jsonToMap map[string]interface{}

		byteCode, _ := json.Marshal(request)

		json.Unmarshal(byteCode, &jsonToMap)

		for key, value := range jsonToMap {
			if value == "" || value == 0 {
				var errorString string = "El campo " + key + " es obligatorio"
				context.JSON(404, gin.H{"Error": errorString})
				return
			}
		}
		//repo
		if len(personas) > 0 {
			for _, persona := range personas {
				if persona.ID > maxId {
					maxId = persona.ID
				}
			}
			request.ID = maxId + 1

		} else {
			request.ID = 1
		}

		personas = append(personas, request) //repo

		context.JSON(200, gin.H{"response": personas})
	}
}

func main() {

	router := gin.Default()

	router.POST("/persona", handlerPersona())

	router.Run()
}
