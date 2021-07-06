package main

import (
	"fmt"
	"io/ioutil"
)

type Product struct {
	id    string
	price string
	total string
}

func save(prod []Product) {
	var text string

	for i, p := range prod {
		text += p.id + "," + p.price + "," + p.total
		if (i + 1) < len(prod) {
			text += ";"
		}
	}
	data := []byte(text)

	err := ioutil.WriteFile("/Users/kcuadrado/meli_bootcamp10/meli_bootcamp10/4_gobases3/tm/ejercicio1.txt", data, 0644)

	if err != nil {
		fmt.Println("ERROR! No se pudo crear el archivo. Descripcion del error: ", err)
	}
}

func main() {
	p1 := Product{id: "111223", price: "30012.00", total: "1"}
	p2 := Product{id: "444321", price: "1000000.00", total: "4"}
	p3 := Product{id: "434321", price: "50.50", total: "1"}

	var productos []Product

	productos = append(productos, p1)
	productos = append(productos, p2)
	productos = append(productos, p3)

	save(productos)
}
