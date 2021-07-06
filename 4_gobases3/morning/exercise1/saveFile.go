package main

import (
	"fmt"
	"io/ioutil"
)

var nextId int = 1

type Product struct {
	id     int
	price  float64
	amount int
}

func newProduct(thePrice float64, theAmount int) Product {
	product := Product{price: thePrice, amount: theAmount}
	product.id = nextId
	nextId++

	return product
}

func (p Product) toString() string {
	return fmt.Sprint(p.id, ";", p.price, ";", p.amount)
}
func main() {
	p1 := newProduct(5000, 4)
	p2 := newProduct(2000, 10)
	p3 := newProduct(30000, 20)

	products := []Product{p1, p2, p3}
	toWrite := ""
	for i, product := range products {
		toWrite += fmt.Sprint(product.toString())
		if i != len(products)-1 {
			toWrite += "\n"
		}
	}
	err := ioutil.WriteFile("./productList.txt", []byte(toWrite), 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success :)")
	}

}
