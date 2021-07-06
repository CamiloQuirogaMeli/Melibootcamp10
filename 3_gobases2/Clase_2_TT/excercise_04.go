package main

import (
	"fmt"
	"math"
)

const (
	MAX_SIZE_SHIPMENT = 10000000.000
)

func main() {

	p1 := Product{size: 5000000.00, typeProduct: "large"}
	p2 := Product{size: 2000000.00, typeProduct: "fragile"}
	p3 := Product{size: 2500.00, typeProduct: "small"}
	p4 := Product{size: 350500.00, typeProduct: "fragile"}

	f := Freight{products: []Product{p1, p2, p3}}
	f.addProduct(p4)

	fmt.Printf("Number of shipments: [%0.f]\n", f.calculateShipment())
}

type Product struct {
	size        float64
	typeProduct string
}

type Freight struct {
	products []Product
}

func (f *Freight) addProduct(p Product) {
	f.products = append(f.products, p)
}

func (f Freight) calculateShipment() float64 {
	totalSize := 0.0
	for _, p := range f.products {
		totalSize += p.totalSize()
	}
	return math.Ceil(totalSize / MAX_SIZE_SHIPMENT)
}

func (p Product) totalSize() float64 {
	switch p.typeProduct {
	case "small":
		return p.size
	case "medium":
		aditional := p.size * 0.05
		return (p.size + aditional)
	case "large":
		aditional := p.size * 0.20
		return (p.size + aditional)
	case "fragile":
		aditional := p.size * 0.75
		return (p.size + aditional)
	}

	return p.size
}
