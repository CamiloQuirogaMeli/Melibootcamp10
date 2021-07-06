package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Product struct {
	id       string  `json:id`
	price    float64 `json:price`
	quantity int     `json:quantity`
}

func main() {

	byteValue, _ := ioutil.ReadFile("./products.json")

	var products []Product

	json.Unmarshal(byteValue, &products)

	fmt.Printf("ID\t Precio\t Cantidad\t")

	for _, product := range products {
		fmt.Printf("%s\t", product.id)
		fmt.Printf("%.2f\t", product.price)
		fmt.Printf("%d\t", product.quantity)
		fmt.Println()
	}
}
