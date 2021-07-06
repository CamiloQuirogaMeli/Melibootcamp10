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
	size     float64
	category string
}

type Flete struct {
	products []Product
}

func (p Product) totalSize() float64 {
	switch p.category {
	case small:
		return p.size
	case medium:
		return p.size * 1.05
	case large:
		return p.size * 1.20
	case fragile:
		return p.size * 1.75
	case special:
		return p.size
	default:
		return 0
	}
}

func (f *Flete) addProduct(p Product) {
	f.products = append(f.products, p)
}

func (f Flete) numberOfShipments() int {
	var total, totalSpecials float64 = 0, 0
	for _, product := range f.products {
		if product.category == special {
			totalSpecials += product.totalSize()
		} else {
			total += product.totalSize()
		}
	}
	if math.Mod(total, 10000000) != 0 {
		total = math.Trunc(total/10000000) + 1
	} else {
		total = math.Trunc(total / 10000000)
	}
	if math.Mod(totalSpecials, 10000000) != 0 {
		totalSpecials = math.Trunc(total/10000000) + 1
	} else {
		totalSpecials = math.Trunc(total / 10000000)
	}
	return int(total + totalSpecials)
}

func main() {
	products1 := []Product{{20000, small}, {1000000, medium}, {1200000, large}, {3000000, large}, {4000000, fragile}}
	flete1 := Flete{
		products: products1,
	}
	products2 := []Product{{1000000, large}, {500000, medium}, {10000, special}}
	flete2 := Flete{
		products: products2,
	}
	flete1.addProduct(Product{5000000, fragile})
	flete2.addProduct(Product{1300000, special})
	fmt.Println("Cantidad de envios flete 1: ", flete1.numberOfShipments())
	fmt.Println("Cantidad de envios flete 2: ", flete2.numberOfShipments())
}
