package main

import (
	"fmt"
)

const (
	SMALL   = "pequeño"
	MEDIUM  = "mediano"
	BIG     = "grande"
	FRAGILE = "fragil"
	SPECIAL = "especial"
)

type Product struct {
	Size float64
	Type string
}

func (product Product) totalSize() float64 {
	switch product.Type {
	case SMALL, SPECIAL:
		return product.Size
	case MEDIUM:
		return product.Size * 1.05
	case BIG:
		return product.Size * 1.2
	case FRAGILE:
		return product.Size * 1.75
	default:
		return product.Size
	}
}

type Flete struct {
	products []Product
}

func (flete *Flete) addProduct(product Product) {
	flete.products = append(flete.products, product)
}

func (flete *Flete) calculatesDeliveries() int {
	max_capacity := 10000000.0
	flete_size := 0.0
	for _, product := range flete.products {
		flete_size += product.totalSize()
	}
	return int(max_capacity/flete_size) + 1
}

func main() {
	p1 := Product{100000, BIG}
	p2 := Product{2300, MEDIUM}
	p3 := Product{450000, SPECIAL}
	p4 := Product{1000000, BIG}
	p5 := Product{123, SMALL}

	flete := Flete{}
	flete.addProduct(p1)
	flete.addProduct(p2)
	flete.addProduct(p3)
	flete.addProduct(p4)
	flete.addProduct(p5)

	fmt.Println("Productos en el flete", flete.products)
	fmt.Println("Cantidad de envios que se deben hacer", flete.calculatesDeliveries())

}
