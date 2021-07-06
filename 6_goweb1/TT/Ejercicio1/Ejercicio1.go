package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	/*
		Ejercicio 1 - Filtremos nuestro endpoint
		Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se
		tiene que poder filtrar por todos los campos.
		1.	Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
		2.	Luego genera la lógica de filtrado de nuestro array.
		3.	Devolver por el endpoint el array filtrado.
	*/

	router := gin.Default()
	router.GET("/users", GetAll)
	router.Run()
}

type User struct {
	Id           int
	Name         string
	Surname      string
	Email        string
	Age          int
	Height       int
	Active       bool
	CreationDate string
}

func GetAll(c *gin.Context) {
	users := []User{}
	dJson, _ := ioutil.ReadFile("users.json")
	err := json.Unmarshal(dJson, &users)
	if err != nil {
		panic("json fail")
	}
	c.JSON(200, users)
}
