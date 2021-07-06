package main

import (
	f "fmt"
	"io/ioutil"
)

type Product struct {
	Id     string
	Nombre string
	Price  float64
	Quantity int
}

func main() {
	var str string

	product := Product{
		Id:     "1",
		Nombre: "Soda",
		Price:  10.99,
		Quantity: 2,
	}
	product2 := Product{
		Id:     "2",
		Nombre: "Chips",
		Price:  5.99,
		Quantity: 5,
	}
	product3 := Product{
		Id:     "3",
		Nombre: "Candy",
		Price:  0.99,
		Quantity: 3,
	}

	products := []Product{product, product2, product3}
	
	for _, p := range products {
		line := f.Sprintf("%s;%s;%.2f;%d\n", p.Id, p.Nombre, p.Price, p.Quantity)
		str += line
	}
	
	data := []byte(str)
	ioutil.WriteFile("./archivo.txt", data, 0644)
}
