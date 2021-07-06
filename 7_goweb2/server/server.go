package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tomastropea/meli_bootcamp10/server/handlers"
	"github.com/tomastropea/meli_bootcamp10/product"
	"encoding/json"
	"io/ioutil"
	"errors"
)

const PRODUCTS_PATH = "data/products.json"
var products []product.Product

const AUTH_TOKEN = "secret token"

func RunServer() {
	var productsMaxID int
	var err error

	if products, err = LoadProductsFromFile(PRODUCTS_PATH); err != nil {
		panic(err) 
	}
	for _, p := range products {
		if productsMaxID < p.ID { 
			productsMaxID = p.ID
		}
	}

	router  := gin.Default()
	
	productsGroup := router.Group("/products")
	{
		productsGroup.GET("/", handlers.GetAllHandler(&products))
		productsGroup.GET("/filter", handlers.FilterProductsHandler(&products))
		productsGroup.GET("/:id", handlers.GetProductByIdHandler(&products))
		productsGroup.POST("/", handlers.CreateProductHandler(AUTH_TOKEN, AddNewProductToDatabase(productsMaxID)))
	}

	router.Run()
}

func AddNewProductToDatabase(productsMaxID int) func (product.Product) (product.Product, error) {

	return func (product product.Product) (product.Product, error) {
		productsMaxID++
		product.ID = productsMaxID
		products = append(products, product)
	
		if err := SaveProductsToDatabase(PRODUCTS_PATH, &products); err != nil {
			return product, err
		}
	
		return product, nil
	}
}

func SaveProductsToDatabase(fileName string, products *[]product.Product) error {
	
	if productsJson, err := json.Marshal(*products); err != nil {
		return err
	} else {
		if err = ioutil.WriteFile(fileName, productsJson, 644); err != nil {
			return err
		}
	}

	return nil
}

func LoadProductsFromFile(fileName string) ([]product.Product, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {

		return make([]product.Product, 0), errors.New("Failed to load products from file")
	} 

	var products []product.Product
	if err := json.Unmarshal(data, &products); err != nil {

		return make([]product.Product, 0), errors.New("Failed to load products from file")
	}

	return products, nil
}