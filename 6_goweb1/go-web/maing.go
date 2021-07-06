// Ejercicio 2 - Hola {nombre}

// Crea dentro de la carpeta go-web un archivo llamado main.go
// Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
// Pegale al endpoint para corroborar que la respuesta sea la correcta.

// Ejercicio 3 - Listar Entidad

// Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
// Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
// Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
// Genera un handler para el endpoint llamado “GetAll”.
// Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	gin "github.com/gin-gonic/gin"
)

type Transaction struct {
	Id              int    `json:"Id"`
	TransactionCode string `json:"TransactionCode"`
	Currency        string `json:"Currency"`
	Sender          string `json:"Sender"`
	Receiver        string `json:"Receiver"`
	Date            string `json:"Date"`
}

func getAll(ctx *gin.Context) {
	trans := []Transaction{}
	content, err := ioutil.ReadFile("./transactions.json")
	if err != nil {
		log.Fatal(err)
	}
	println(string(content))
	errJ := json.Unmarshal([]byte(content), &trans)
	if errJ != nil {
		log.Fatal(err)
	}
	ctx.JSON(200, gin.H{
		"Message": trans,
	})
}

func main() {

	router := gin.Default()

	router.GET("/transactions", getAll)
	router.Run()
}
