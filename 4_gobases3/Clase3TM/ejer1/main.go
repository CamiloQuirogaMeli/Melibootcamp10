package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	p := []byte(`[
		{"id": 100, "price": 500, "quantity": 1}, 
		{"id": 200, "price": 1500, "quantity": 5}, 
		{"id": 300, "price": 2500, "quantity": 10}
		]`)
	err := ioutil.WriteFile("./products.json", p, 0644)

	if err != nil {
		fmt.Print(err)
	}
}
