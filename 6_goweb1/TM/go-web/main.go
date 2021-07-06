package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	IsPublished  bool    `json:"isPublished"`
	CreationDate string  `json:"creationDate"`
}

func GetAll(c *gin.Context) {
	products := []Product{}
	byteJson, _ := ioutil.ReadFile("products.json")
	err := json.Unmarshal(byteJson, &products)
	if err != nil {
		panic("unparseable json")
	}
	c.JSON(200, products)
}

func main() {
	router := gin.Default()

	router.GET("/saludo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Santiago!",
		})
	})

	router.GET("/products", GetAll)

	router.Run(":8083")

}
