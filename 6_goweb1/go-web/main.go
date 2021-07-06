package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

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

func main() {

	// product := Product{Id: 1, Name: "Andres", Color: "Rojo"}
	// datajson, err := json.Marshal(product)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(string(datajson))
	// }
	fmt.Println("server starting...")
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola",
		})
	})
	router.GET("/fake-products", GetMocked)
	products := router.Group("/products")
	{
		products.GET("/", GetAll)
		products.GET("/:id", GetById)
		products.GET("/by", GetBy)
	}

	router.Run()

}

func GetAll(c *gin.Context) {
	byteProducts, err := ioutil.ReadFile("./products.json")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		log.Fatal("Can't read json file:", err)
	} else {
		var products []Product
		if err := json.Unmarshal(byteProducts, &products); err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			log.Fatal(err)
		} else {
			c.JSON(200, gin.H{
				"data": products,
			})
			fmt.Println(products)
		}

	}
}
func GetBy(c *gin.Context) {
	byteProducts, err := ioutil.ReadFile("./products.json")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		log.Fatal("Can't read json file:", err)
	} else {
		var products []Product
		if err := json.Unmarshal(byteProducts, &products); err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			log.Fatal(err)
		} else {

			var filteredProducts []Product
			match := false
			for _, product := range products {
				if c.Query("id") == fmt.Sprintf("%d", product.Id) {
					match = true
				} else if c.Query("id") != "" {
					match = false
				}
				if c.Query("name") == product.Name && match {
					match = true
				} else if c.Query("name") != "" {
					match = false
				}
				if c.Query("color") == product.Color && match {
					match = true
				} else if c.Query("color") != "" {
					match = false
				}
				if c.Query("price") == fmt.Sprintf("%d", product.Price) && match {
					match = true
				} else if c.Query("price") != "" {
					match = false
				}
				if c.Query("stock") == fmt.Sprintf("%d", product.Stock) && match {
					match = true
				} else if c.Query("stock") != "" {
					match = false
				}
				if c.Query("creation-date") == product.CreationDate && match {
					match = true
				} else if c.Query("creation-date") != "" {
					match = false
				}
				if match {
					filteredProducts = append(filteredProducts, product)
				}

			}
			c.JSON(200, gin.H{
				"data": filteredProducts,
			})
			fmt.Println(products)
		}

	}
}
func GetById(c *gin.Context) {
	byteProducts, err := ioutil.ReadFile("./products.json")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		log.Fatal("Can't read json file:", err)
	} else {
		var products []Product
		if err := json.Unmarshal(byteProducts, &products); err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			log.Fatal(err)
		} else {
			match := false
			for _, product := range products {
				if c.Param("id") == fmt.Sprintf("%d", product.Id) {
					match = true
					c.JSON(200, gin.H{
						"data": product,
					})
				}
			}
			if !match {
				c.JSON(404, gin.H{
					"error": "Product not found",
				})
			}
		}

	}
}
func GetMocked(c *gin.Context) {

	products := []Product{
		{
			Id:    1,
			Name:  "Andres",
			Color: "Rojo",
		},
		{
			Id:    2,
			Name:  "Felipe",
			Color: "Azul",
		}}

	c.JSON(200, gin.H{
		"data": products,
	})

	// datajson, err := json.Marshal(products)
	// if err != nil {
	// 	c.JSON(500, gin.H{
	// 		"Error": err,
	// 	})
	// } else {
	// 	fmt.Println(string(datajson))
	// 	c.JSON(200, gin.H{
	// 		"data": datajson,
	// 	})
	// }

}
