package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

var products = []Product{}

type Product struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      string  `json:"code"`
	Published bool    `json:"published"`
	CreatedAt string  `json:"createdAt"`
}

func main() {
	GetProducts()
	router := gin.Default()
	router.GET("/saludo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Alejandro",
		})
	})

	product := router.Group("/products")
	{
		product.GET("/", HandlerGetAll)
		product.GET("/:id", HandlerGetById)
		product.POST("/", HandlerAddProduct)
	}

	router.Run()
}

func HandlerGetAll(c *gin.Context) {
	products, err := GetProducts()

	products = filterProducts(&products, "id", c.Query("id"))
	products = filterProducts(&products, "name", c.Query("name"))
	products = filterProducts(&products, "color", c.Query("color"))
	products = filterProducts(&products, "price", c.Query("price"))
	products = filterProducts(&products, "stock", c.Query("stock"))
	products = filterProducts(&products, "createdAt", c.Query("createdAt"))
	products = filterProducts(&products, "code", c.Query("code"))
	products = filterProducts(&products, "published", c.Query("published"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	if len(products) == 0 {
		c.JSON(404, gin.H{"error": "products not found"})
		return
	}
	c.JSON(200, gin.H{"products": products})
}

func filterProducts(productList *[]Product, filter string, value string) []Product {
	if value == "" {
		return *productList
	}
	var filterList = []Product{}
	for _, product := range *productList {
		switch filter {
		case "id":
			if fmt.Sprint(product.Id) == value {
				filterList = append(filterList, product)
			}
		case "name":
			if product.Name == value {
				filterList = append(filterList, product)
			}
		case "color":
			if product.Color == value {
				filterList = append(filterList, product)
			}
		case "price":
			if fmt.Sprint(product.Price) == value {
				filterList = append(filterList, product)
			}
		case "stock":
			if fmt.Sprint(product.Stock) == value {
				filterList = append(filterList, product)
			}
		case "code":
			if product.Code == value {
				filterList = append(filterList, product)
			}
		case "published":
			if fmt.Sprint(product.Published) == value {
				filterList = append(filterList, product)
			}
		case "createdAt":
			if product.CreatedAt == value {
				filterList = append(filterList, product)
			}
		}

	}
	return filterList
}

func HandlerGetById(c *gin.Context) {
	products, err := GetProducts()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	for _, prod := range products {
		if fmt.Sprint(prod.Id) == c.Param("id") {
			c.JSON(200, prod)
			return
		}
	}

	c.JSON(404, gin.H{
		"error": "Product not found",
	})
}

func GetProducts() ([]Product, error) {
	if len(products) > 0 {
		return products, nil
	}
	productsJSON, err := ioutil.ReadFile("./products.json")
	if err != nil {
		log.Fatal(err)
	}
	// products := []Product{}
	err = json.Unmarshal(productsJSON, &products)
	if err != nil {
		log.Fatal(err)
	}
	return products, err
}

// func Clase 2 TM
func HandlerAddProduct(c *gin.Context) {
	_, err := ValidateToken(c.Request.Header.Get("token"))
	if err != nil {
		c.JSON(401, gin.H{
			"error": fmt.Sprint(err),
		})
		return
	}
	var req Product
	c.Bind(&req)
	_, errProd := ValidateProduct(&req)
	if errProd != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprint(errProd),
		})
		return
	}
	req.Id = len(products) + 1
	products = append(products, req)
	c.JSON(200, req)
}

func ValidateProduct(prod *Product) (bool, error) {
	switch {
	case prod.Name == "":
		return false, errors.New("el campo name es requerido")
	case prod.Color == "":
		return false, errors.New("el campo color es requerido")
	case prod.Price == 0:
		return false, errors.New("el campo price es requerido")
	case prod.Stock == 0:
		return false, errors.New("el campo stock es requerido")
	case prod.Code == "":
		return false, errors.New("el campo code es requerido")
	case !prod.Published:
		return false, errors.New("el campo published es requerido")
	case prod.CreatedAt == "":
		return false, errors.New("el campo createdAt es requerido")
	default:
		return true, nil
	}
}
