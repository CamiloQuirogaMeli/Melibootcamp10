package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published"`
	CreationDate string `json:"creation_date"`
}

var products []Product

const TOKEN = "123"

func main() {

	fmt.Println("server starting...")
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola",
		})
	})
	products := router.Group("/products")
	{
		products.POST("/new", NewProduct())
	}

	router.Run()

}
func GetNextId() int {
	if len(products) == 0 {
		return 1
	} else {
		return products[len(products)-1].Id + 1
	}
}

func checkProduct(product Product) (bool, string) {
	s := reflect.ValueOf(&product).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		v := reflect.ValueOf(f.Interface())
		if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
			return false, fmt.Sprint(s.Type().Field(i).Tag.Get("json"))
		}
	}
	return true, ""
}

func NewProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token := c.Request.Header.Get("token"); token != TOKEN {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is not valid",
			})
			return
		}
		var req Product
		if err := c.Bind(&req); err != nil {
			c.JSON(403, gin.H{
				"error": err,
			})
			return
		}
		req.Id = GetNextId()
		if valid, fieldName := checkProduct(req); !valid {
			c.JSON(400, gin.H{
				"error": fmt.Sprintf("Missing %s fieldName in request body", fieldName),
			})
			return
		}
		products = append(products, req)

		c.JSON(http.StatusCreated, req)
		fmt.Println(products)

	}
}
