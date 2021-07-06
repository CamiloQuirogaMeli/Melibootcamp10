package main

import (
	"fmt"
	"math"
)

const (
	small   = "small"
	medium  = "medium"
	large   = "large"
	fragile = "fragile"
	special = "special"
)

type Product struct {
	prodType string
	size     float64
}

type Freight struct {
	products []Product
}

func (f *Freight) AddProduct(p Product) {
	f.products = append(f.products, p)
}

// NumOfShipments calculates the number of shipments
// that the Freight requieres.
// TODO: Need to add logic for special products
func (f *Freight) NumOfShipments() float64 {
	var toShip float64
	var shipments float64
	for _, product := range f.products {
		toShip += product.size
	}
	shipments = math.Ceil(toShip / 10000000)
	return shipments
}

func (p *Product) TotalSize() float64 {
	switch p.prodType {
	case medium:
		return p.size + (0.5 * p.size)
	case large:
		return p.size + (0.20 * p.size)
	case fragile:
		return p.size + (0.75 * p.size)
	}
	return p.size
}

func main() {
	prod1 := Product{
		prodType: fragile,
		size:     10000000,
	}
	prod2 := Product{
		prodType: fragile,
		size:     10000000,
	}
	products := []Product{prod1, prod2}
	freight := Freight{
		products: products,
	}
	prod3 := Product{
		prodType: large,
		size:     30000000,
	}
	freight.AddProduct(prod3)
	fmt.Println(freight.NumOfShipments())
}
