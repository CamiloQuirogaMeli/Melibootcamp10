package main

import (
	"fmt"
)

const (
	SMALL   = "small"
	MEDIUM  = "medium"
	LARGE   = "large"
	FRAGILE = "fragile"
	SPECIAL = "special"
)

type Product interface {
	totalSize() float64
}

type Truck interface {
	addProduct(Product)
	calculateShippings() int
}

type Small struct {
	size float64
}

type Medium struct {
	size float64
}

type Large struct {
	size float64
}

type Fragile struct {
	size float64
}

type Special struct {
	size float64
}

type MeliTruck struct {
	products []Product
	capacity float64
}

func (s Small) totalSize() float64 {
	return s.size
}

func (m Medium) totalSize() float64 {
	return m.size * (1 + 0.05)
}

func (l Large) totalSize() float64 {
	return l.size * (1 + 0.2)
}

func (f Fragile) totalSize() float64 {
	return f.size * (1 + 0.75)
}

func (sp Special) totalSize() float64 {
	return sp.size
}

func (mt *MeliTruck) addProduct(product Product) {
	mt.products = append(mt.products, product)
}

func createProduct(productType string, productSize float64) Product {
	switch productType {
	case SMALL:
		return Small{size: productSize}
	case MEDIUM:
		return Medium{size: productSize}
	case LARGE:
		return Large{size: productSize}
	case FRAGILE:
		return Fragile{size: productSize}
	case SPECIAL:
		return Special{size: productSize}
	}
	return nil
}

func (mt MeliTruck) calculateShippings() int {
	actualCapacity, shippings := 0.0, 0
	specialProducts := 0
	for _, product := range mt.products {
		if actualCapacity+product.totalSize() < mt.capacity {
			switch product.(type) {
			case Special:
				specialProducts++
			default:
				shippings++
			}
			actualCapacity += product.totalSize()
		} else {
			break
		}
	}
	if specialProducts > 0 {
		return specialProducts
	} else {
		return shippings
	}
}

func main() {
	meliTruck := MeliTruck{capacity: 10000000}

	smallProduct := Small{size: 30}
	meliTruck.addProduct(smallProduct)
	mediumProduct := Medium{size: 60}
	meliTruck.addProduct(mediumProduct)
	specialProduct := Special{size: 60}
	meliTruck.addProduct(specialProduct)

	fmt.Println(meliTruck.products)
	fmt.Println(meliTruck.calculateShippings())
}
