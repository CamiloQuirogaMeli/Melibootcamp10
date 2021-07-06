package main

import (
	f "fmt"
)

const (
	CHICO     = "chico"
	MEDIANO   = "mediano"
	GRANDE    = "grande"
	FRAGIL    = "fragil"
	ESPECIAL  = "especial"
	MAX_SPACE = 10000000
)

type product struct {
	typeProduct string
	size        float32
}

type freight struct {
	products []product
	space    float32
	special  bool
}

func (freight *freight) addProduct(product product) {
	freight.products = append(freight.products, product)
}
func (freight freight) calcShipment() int {
	return len(freight.products)
}

func (product product) totalSize() float32 {
	switch product.typeProduct {
	case CHICO:
		return product.size
	case MEDIANO:
		return product.size * 1.05
	case GRANDE:
		return product.size * 1.20
	case FRAGIL:
		return product.size * 1.75
	default:
		return 0
	}
}

func main() {
	var flag bool = false
	var products []product

	freight := freight{space: MAX_SPACE}
	products = append(products, product{typeProduct: MEDIANO, size: 500000})
	products = append(products, product{typeProduct: ESPECIAL, size: 100000})
	products = append(products, product{typeProduct: GRANDE, size: 1000000})
	products = append(products, product{typeProduct: GRANDE, size: 1000000})
	products = append(products, product{typeProduct: ESPECIAL, size: 100000})
	products = append(products, product{typeProduct: GRANDE, size: 1000000})
	products = append(products, product{typeProduct: GRANDE, size: 1000000})
	products = append(products, product{typeProduct: GRANDE, size: 1000000})
	products = append(products, product{typeProduct: GRANDE, size: 1000000})
	products = append(products, product{typeProduct: GRANDE, size: 9000000})
	products = append(products, product{typeProduct: GRANDE, size: 1000000})
	products = append(products, product{typeProduct: FRAGIL, size: 1000000})
	products = append(products, product{typeProduct: ESPECIAL, size: 100000})

	if products[0].typeProduct == ESPECIAL {
		freight.special = true
	} else {
		freight.special = false
	}
	for _, product := range products {
		f.Println(freight.special)
		if freight.space > product.totalSize() && !flag {
			if freight.special && product.typeProduct == ESPECIAL {
				freight.space -= product.size
				freight.addProduct(product)
			} else if !freight.special && product.typeProduct != ESPECIAL{
				freight.space -= product.size
				freight.addProduct(product)
				
			}
		} else {
			flag = !flag
		}
	}

	f.Printf("Total shipments %d\n", freight.calcShipment())
	f.Println("Freight details: ")
	for i, product := range freight.products {
		f.Printf("Product: %d Type: %s Size %.2f cm3\n", i+1, product.typeProduct, product.size)
	}
}
