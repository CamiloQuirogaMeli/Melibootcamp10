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

func getById(c *gin.Context) {
	p := []producto{}
	content, err := ioutil.ReadFile("../../products.json")
	if err != nil {
		log.Fatal(err)
	}

	err1 := json.Unmarshal(content, &p)

	if err1 != nil {
		log.Fatal(err1)
	}

	for _, prod := range p {
		id := c.Param("id")
		if id == strconv.Itoa(prod.Id) {
			c.JSON(200, gin.H{
				"Message": prod,
			})
			break
		}
	}

	c.JSON(404, gin.H{
		"Message": "Item not found",
	})

}

func main() {
	r := gin.Default()
	r.GET("/productos/:id", getById)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
