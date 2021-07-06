package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Prod struct {
	Id            int     `form: "id" json:"id"`
	Nombre        string  `form: "nombre" json:"nombre"`
	Color         string  `form: "color" json:"color"`
	Precio        float32 `form: "precio" json:"precio"`
	Stock         int     `form: "stock" json:"stock"`
	Codigo        string  `form: "codigo" json:"codigo"`
	Publicado     bool    `form: "publicado" json:"publicado"`
	FechaCreacion string  `form: "fecha_creacion" json:"fecha_creacion"`
}

func index(ctxt *gin.Context) {
	ctxt.String(200, "Pagina principal")
}

func getAll(ctxt *gin.Context) {
	archivoJson := leerArchivoJson()
	//ctxt.String(200, string(archivoJson))
	p := []Prod{}
	err := json.Unmarshal([]byte(archivoJson), &p)
	if err != nil {
		log.Fatal()
	}

	ctxt.JSON(200, p)
}

func leerArchivoJson() []byte {
	prod, err1 := ioutil.ReadFile("./products.json")
	if err1 != nil {
		log.Fatal()
	}
	return prod
}

func getPorId(ctxt *gin.Context) {
	archivoJson := leerArchivoJson()

	p := []Prod{}
	err := json.Unmarshal([]byte(archivoJson), &p)
	if err != nil {
		log.Fatal()
	}

	valor, _ := strconv.Atoi(ctxt.Param("id"))
	var existe bool = false
	for _, h := range p {
		if h.Id == valor {
			existe = true
			break
		}
	}

	if existe {
		ctxt.String(200, "El empleado: %s existe.", ctxt.Param("id"))
	} else {
		ctxt.String(404, "El empleado: %s NO existe.", ctxt.Param("id"))
	}

}

func getByColor(ctxt *gin.Context) {
	archivoJson := leerArchivoJson()

	listaProductos := []Prod{}
	filtrados := []Prod{}
	var cantidad int

	err := json.Unmarshal([]byte(archivoJson), &listaProductos)
	if err != nil {
		log.Fatal()
	}

	for _, p := range listaProductos {
		if ctxt.Query("color") == p.Color {
			filtrados = append(filtrados, p)
			cantidad++
		}
	}

	if cantidad > 0 {
		ctxt.String(200, "Filtrado exitoso")
		for _, p := range filtrados {
			fmt.Println("Id: ", p.Id)
			fmt.Println("Nombre: ", p.Nombre)
			fmt.Println("Color: ", p.Color)
		}
	} else {
		ctxt.String(404, "No se ha podido filtrar")
	}

}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hola Santiago!"}) })
	router.GET("/index", index)
	router.GET("/products", getAll)
	router.GET("/products/:id", getPorId)
	router.GET("/productsByColor", getByColor)
	router.Run()
}
