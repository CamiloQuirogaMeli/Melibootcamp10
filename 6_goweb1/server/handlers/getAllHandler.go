package handlers

import (
	"encoding/json"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/tomastropea/meli_bootcamp10/product"
	"errors"
)

func GetAllHandler(ctx *gin.Context) {

	if products, err := LoadProductsFromFile(PRODUCTS_PATH); err != nil {

		ctx.JSON(500, gin.H{"error": err}) 
		return
	} else {

		ctx.JSON(200, gin.H{
			"product.products": products,
		}) 
	}
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