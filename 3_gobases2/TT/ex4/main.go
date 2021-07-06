package main

import (
	"fmt"
)

const (
	SMALL       = "small"
	MEDIUM      = "medium"
	LARGE       = "large"
	FRAGILE     = "fragile"
	SPECIAL     = "special"
	MAXCAPACITY = 10000000
)

type product struct {
	Id       int
	Name     string
	sizeType string
	size     float64
	price    float64
}

func (p *product) totalSize() float64 {
	switch p.sizeType {
	case SMALL:
		return p.size
	case MEDIUM:
		return p.size * 1.05
	case LARGE:
		return p.size * 1.2
	case FRAGILE:
		return p.size * 1.75
	case SPECIAL:
		return p.size
	}
	return 0
}

type cargo struct {
	products []product
}

func (c *cargo) addProduct(id int, name string, sizeType string, size float64, price float64) {
	c.products = append(c.products, product{id, name, sizeType, size, price})
}

func (c cargo) calcShipping() {
	totalSize := 0.0
	totalSpecialSize := 0.0
	for _, p := range c.products {
		if p.sizeType != SPECIAL {
			totalSize += p.totalSize()
		} else {
			totalSpecialSize += p.totalSize()
		}
	}
	shipping := (MAXCAPACITY / totalSize)
	specialShipping := (MAXCAPACITY / totalSpecialSize)
	fmt.Println("Cantidad de envíos para productos no especiales: ", shipping)
	fmt.Println("Cantidad de envíos para productos especiales: ", specialShipping)
}

func main() {
	cargo := cargo{}
	cargo.addProduct(1, "Cepillo", SMALL, 1000, 4000)
	cargo.calcShipping()
}
