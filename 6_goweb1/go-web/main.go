package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Transactions struct {
	Id       string
	Code     string
	Currency string
	Amount   float64
	Sender   string
	Receiver string
	Date     string
}

func main() {
	fmt.Println("Hola")
	// Creates a gin router with default middleware
	router := gin.Default()
	// A handler for GET request on /example

	// router.GET("/hello-world", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World!",
	// 	}) // gin.H is a shortcut for map[string]interface{}
	// })

	router.GET("/hello", getName)
	router.GET("/transactions", getTransactions)

	router.Run() // listen and serve on port 8080 }
}

func getName(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	}) // gin.H is a shortcut for map[string]interface{}
}

func getTransactions(c *gin.Context) {
	data, err := ioutil.ReadFile("./transactions.json")

	if err != nil {
		log.Fatal((err))
		os.Exit(1)
	}
	t := []Transactions{}

	if err := json.Unmarshal([]byte(data), &t); err != nil {
		log.Fatal(err)
	}
	c.JSON(200, t)
}

//Esta función que realizamos nos permite ver la anatomía de un mensaje Request de una
//forma más práctica.
func Ejemplo(context *gin.Context) {
	//El body, header y method están contenidos en el contexto de gin.
	contenido := context.Request.Body
	header := context.Request.Header
	metodo := context.Request.Method
	fmt.Println("¡He recibido algo!")
	fmt.Printf("\tMetodo: %s\n", metodo)
	fmt.Printf("\tContenido:\n")
	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	}
	fmt.Printf("\tCotenido:%s\n", contenido)
	fmt.Println("¡Yay!")
	context.String(200, "¡Lo recibí!") //Respondemos al cliente con 200 OK y un bonito
 mensaje.
 }