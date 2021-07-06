package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

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

func getAll(c *gin.Context) {
	p := []producto{}
	content, err := ioutil.ReadFile("../products.json")
	if err != nil {
		log.Fatal(err)
	}

	err1 := json.Unmarshal(content, &p)

	if err1 != nil {
		log.Fatal(err1)
	}

	c.JSON(200, gin.H{
		"Message": p,
	})

}

func main() {
	r := gin.Default()
	/*r.GET("/hola-adrian", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "Hola Adrian",
		})
	})*/
	r.GET("/productos", getAll)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
