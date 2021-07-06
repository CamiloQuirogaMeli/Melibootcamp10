package main

import (
	"fmt"
	"io/ioutil"
)

type product struct {
	id       int
	price    float64
	quantity int
}

func main() {

	prod1 := product{1, 10.00, 5}
	prod2 := product{2, 100.00, 10}
	prod3 := product{3, 20.00, 3}

	products := []product{prod1, prod2, prod3}

	var prodStr string

	for _, prod := range products {

		prodStr += fmt.Sprintf("%v \t %v \t %v;\n", prod.id, prod.price, prod.quantity)

	}

	byteArray := []byte(prodStr)

	err := ioutil.WriteFile("./file.txt", byteArray, 0644)

	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("File saved.")
	}

}
