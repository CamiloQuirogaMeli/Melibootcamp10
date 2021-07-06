package main

import (
	"fmt"
	"io/ioutil"
)

type Product struct {
	id       int
	price    float32
	quantity int
}

func main() {
	products := []Product{
		{id: 1, price: 100, quantity: 2},
		{id: 2, price: 200, quantity: 4},
		{id: 3, price: 300, quantity: 7},
		{id: 4, price: 400, quantity: 8},
	}
	var text string

	for _, product := range products {
		text += fmt.Sprintf("%d;%.2f;%d\n", product.id, product.price, product.quantity)
	}
	textFile := []byte(text)
	err := ioutil.WriteFile("./go-file-products.txt", textFile, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
