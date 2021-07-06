package main

import (
	"fmt"
	"sort"
)

const (
	SMALL   = "small"
	MEDIUM  = "medium"
	BIG     = "big"
	FRAGILE = "fragile"
	SPECIAL = "special"
)

const (
	CARRIER_LIMIT = 10000000
)

type IProduct interface {
	totalSize()
}

type Product struct {
	size        float64
	productType string
}

type ICarrier interface {
	addProducts()
	calculateProducts()
}

type Carrier struct {
	products []Product
}

func (p Product) totalSize() (float64, int) {

	switch p.productType {
	case MEDIUM:
		return p.size + (p.size * 0.05), 0
	case BIG:
		return p.size + (p.size * 0.2), 0
	case FRAGILE:
		return p.size + (p.size * 0.75), 0
	case SPECIAL:
		return 0, 1
	}

	return p.size, 0
}

func (c *Carrier) addProducts(products []Product) {
	c.products = products
}

func (c Carrier) calculateProducts() {

	var productsToDeliver int
	var carrierTotalSize float64
	var productsSizesOrdered []float64
	var specialProducts int

	for _, product := range c.products {
		totalSize, specialProduct := product.totalSize()
		productsSizesOrdered = append(productsSizesOrdered, totalSize)
		specialProducts += specialProduct
	}

	sort.Float64s(productsSizesOrdered)

	for _, productSize := range productsSizesOrdered {

		totalSize := carrierTotalSize + productSize

		if totalSize >= CARRIER_LIMIT {
			break
		}

		productsToDeliver++
		carrierTotalSize = totalSize
	}

	fmt.Print("The quantity of shippings to deliver is: ", carrierTotalSize, " And the special products to deliver ", specialProducts)
}

func main() {

	product1 := Product{3000, SMALL}
	product2 := Product{3000, MEDIUM}
	product3 := Product{3000, BIG}
	product4 := Product{3000, FRAGILE}
	product5 := Product{3000, SPECIAL}

	carrier1 := Carrier{}

	products := [...]Product{product1, product2, product3, product4, product5}
	carrier1.addProducts(products[:])
	carrier1.calculateProducts()
}
