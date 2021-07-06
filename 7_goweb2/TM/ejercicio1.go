package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Edad     float64 `json:"edad"`
}

var personas []Persona = []Persona{}

func handlerPersona() gin.HandlerFunc {

	return func(context *gin.Context) {

		token := context.Request.Header.Get("token")

		if token != "12345" {
			context.JSON(404, gin.H{
				"error": "No tiene permisos para realizar la peticion",
			})
			return
		}

		var request Persona

		var maxId int = 0

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

		personas = append(personas, request)

		context.JSON(200, gin.H{"response": personas})
	}
}

func main() {

	router := gin.Default()

	router.POST("/persona", handlerPersona())

	router.Run()
}
