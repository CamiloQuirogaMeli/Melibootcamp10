package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	/*
		Ejercicio 2 - Get one endpoint
		Generar un nuevo endpoint que nos permita traer un solo resultado del array de la temática.
		Utilizando path parameters el endpoint debería ser /temática/:id
		(recuerda que siempre tiene que ser en plural la temática). Una vez recibido el id devuelve la posición correspondiente.
		Genera una nueva ruta.
		Genera un handler para la ruta creada.
		Dentro del handler busca el item que necesitas.
		Devuelve el item según el id.
		Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.

	*/

	router := gin.Default()
	router.GET("/users/:Id", GetById)
	router.Run()
}

type User struct {
	Id           string
	Name         string
	Surname      string
	Email        string
	Age          string
	Height       string
	Active       bool
	CreationDate string
}

func GetById(c *gin.Context) {
	users := []User{}
	dJson, err := ioutil.ReadFile("users.json")
	if err != nil {
		panic("json fail")
	} else {
		if err1 := json.Unmarshal(dJson, &users); err1 != nil {
			panic(err1)
		}
	}

	for _, value := range users {
		if value.Id == c.Param("Id") {

			c.JSON(200, gin.H{
				"message": value,
			})
			break
		}
	}

	c.JSON(404, gin.H{
		"Message": "No se encontro el usuario",
	})
}
