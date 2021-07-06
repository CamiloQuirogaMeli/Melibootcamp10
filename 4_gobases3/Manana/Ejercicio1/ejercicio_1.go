package main

import (
	"fmt"
	"io/ioutil"
)

type producto struct {
	id       int
	precio   float64
	cantidad int
}

func getDetail(prods []producto) string {
	var detail string
	for i, prod := range prods {
		if i == 0 {
			detail = fmt.Sprintf("%d %.2f %d; ", prod.id, prod.precio, prod.cantidad)
		} else {
			detail += fmt.Sprintf("%d %.2f %d; ", prod.id, prod.precio, prod.cantidad)
		}
	}
	return detail
}

func main() {
	p1 := producto{1, 15.50, 4}
	p2 := producto{2, 2.74, 100}
	p3 := producto{3, 2370.50, 1}
	p4 := producto{4, 150.0, 11}

	p := []producto{p1, p2, p3, p4}

	d1 := []byte(getDetail(p))
	ioutil.WriteFile("./productos.txt", d1, 0644)

}
