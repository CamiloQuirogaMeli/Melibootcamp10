package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/josephcasa/meli_bootcamp10/products"
)

func GetAll(c *gin.Context) {

	if products, err := LoadFile("products.json"); err != nil {

		c.JSON(500, gin.H{"error": err})
		return

	} else {

		c.JSON(200, gin.H{
			"products": products,
		})
	}
}

func LoadFile(file_name string) ([]products.Product, error) {

	data, err := ioutil.ReadFile(file_name)
	products := []products.Product{}

	if err != nil {
		return products, errors.New("Failed to load products from file")
	}

	if err := json.Unmarshal(data, &products); err != nil {
		return products, errors.New("Failed to load products from file")
	}

	return products, nil
}
