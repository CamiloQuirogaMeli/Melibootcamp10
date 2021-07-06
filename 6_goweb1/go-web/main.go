package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	OnSale       string  `json:"published"`
	CreationDate string  `json:"creation_date"`
}

const (
	PRICE         = "price"
	NAME          = "name"
	COLOR         = "color"
	ON_SALE       = "published"
	CREATION_DATE = "creation_date"
	CODE          = "code"
	ID            = "id"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Conrado!",
		})
	})
	products := r.Group("/products")

	products.GET("", GetAll)
	products.GET("/:id", GetById)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetAll(ctx *gin.Context) {
	products, _ := loadData()

	var response []Product
	price, _ := strconv.ParseFloat(ctx.Query(PRICE), 64)

	params := ctx.Request.URL.Query()

	if len(params) != 0 {
		for _, product := range products {
			if ctx.Query(NAME) == product.Name ||
				price == product.Price || ctx.Query(ON_SALE) == product.OnSale ||
				ctx.Query(CREATION_DATE) == product.CreationDate ||
				ctx.Query(CODE) == product.Code || ctx.Query(COLOR) == product.Color {
				response = append(response, product)
			}
		}
	} else {
		response = products
	}
	ctx.JSON(200, gin.H{
		"message": response,
	})
}

func GetById(ctx *gin.Context) {
	products, _ := loadData()
	for _, p := range products {
		if p.Id == ctx.Param(ID) {
			ctx.JSON(200, gin.H{
				"message": p,
			})
			return
		}
	}
	ctx.JSON(404, gin.H{
		"message": "No se encontro el producto",
	})
}

func loadData() ([]Product, error) {
	data, err := ioutil.ReadFile("./products.json")

	if err != nil {
		return nil, err
	}
	var p []Product
	if err := json.Unmarshal(data, &p); err != nil {
		log.Fatal(err)
	}
	return p, nil
}
