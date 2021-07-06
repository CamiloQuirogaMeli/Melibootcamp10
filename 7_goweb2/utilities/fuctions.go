package utilities

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/josephcasa/meli_bootcamp10/product"
)

func LoadFile(file_name string) ([]product.Product, error) {

	data, err := ioutil.ReadFile(file_name)
	products := []product.Product{}

	if err != nil {
		return products, errors.New("Failed to load products from file")
	}

	if err := json.Unmarshal(data, &products); err != nil {
		return products, errors.New("Failed to load products from file")
	}

	return products, nil
}

func AddNewProduct(products []product.Product, maxId int) func(product.Product) (product.Product, error) {

	return func(product product.Product) (product.Product, error) {
		maxId++
		product.Id = maxId
		products = append(products, product)

		if err := SaveProducts("products.json", &products); err != nil {
			return product, err
		}

		return product, nil
	}
}

func SaveProducts(file_name string, products *[]product.Product) error {

	if productsJson, err := json.Marshal(*products); err != nil {
		return err
	} else if err = ioutil.WriteFile(file_name, productsJson, 404); err != nil {
		return err
	}

	return nil
}
