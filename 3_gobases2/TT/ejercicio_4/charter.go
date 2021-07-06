package main

import (
	product "github.com/tomastropea/TT/ejercicio_4/product"
	"math"
)

type Charter struct {
	items []product.Product
}

func (c *Charter) AddProduct(item product.Product){
	c.items = append(c.items, item)
}

func (c *Charter) RequiredShippings() int {
	const CAPACITY = 10000000.0
	var specialProductsTotalSpace, nonSpecialProductsTotalSpace int

	for _, item := range c.items {
		if item.IsSpecial() {
			specialProductsTotalSpace += item.TotalSize()
		} else {
			nonSpecialProductsTotalSpace += item.TotalSize()
		}
	}

	specialProductsRequiredShippings := math.Ceil(float64(specialProductsTotalSpace) / CAPACITY)
	nonSpecialProductsRequiredShippings := math.Ceil(float64(nonSpecialProductsTotalSpace) / CAPACITY)
	
	return int(specialProductsRequiredShippings + nonSpecialProductsRequiredShippings)
}
