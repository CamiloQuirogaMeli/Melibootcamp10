package main

import (
	"fmt"
	"io/ioutil"
)

type Product struct {
	Id int
	Price float64
	Quantity int
}

func AsString(product Product) string {
	return fmt.Sprintf("%d;%f;%d\n", product.Id, product.Price, product.Quantity)
}

func StoreProducts(products []Product) {

	lines := ""

	for _, product := range products {
		lines += AsString(product)
	}

	ioutil.WriteFile("products.txt", []byte(lines), 0644)
}

func main() {

	products := []Product{
		{111223, 30012.00, 1},
		{444321, 1000000.00, 4},
		{434321, 50.50, 1},
	}

	StoreProducts(products)
}