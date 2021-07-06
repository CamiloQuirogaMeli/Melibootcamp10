package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Prod struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float32 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_creacion"`
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

	// for _, p := range p {
	// 	ctxt.String(200, "Codigo: ", p.Id)
	// 	//fmt.Println("Codigo: ", p.Codigo)
	// }

	ctxt.JSON(200, p)
}

func leerArchivoJson() []byte {
	prod, err1 := ioutil.ReadFile("./products.json")
	if err1 != nil {
		log.Fatal()
	}
	return prod
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hola Santiago!"}) })
	router.GET("/index", index)
	router.GET("/products", getAll)
	router.Run()
}
