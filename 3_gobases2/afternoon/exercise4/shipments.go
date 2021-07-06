package main

import (
	"errors"
	"fmt"
)

const (
	SMALL   = "small"
	MEDIUM  = "medium"
	BIG     = "big"
	FRAGILE = "fragile"
	SPECIAL = "special"

	SHIPPING_LIMIT = 10000000.0
)

type Product struct {
	name        string
	productType string
	size        float64
	space       float64
}

func newProduct(theName string, theProductType string, theSize float64) (Product, error) {

	product := Product{}

	aditionalPercentageSpace := map[string]float64{SMALL: 0, MEDIUM: 0.05, BIG: 0.2, FRAGILE: 0.75, SPECIAL: 0}

	if _, ok := aditionalPercentageSpace[theProductType]; !ok {
		return product, errors.New(theProductType + " is not a valid product type")
	}
	product.name = theName
	product.size = theSize
	product.productType = theProductType
	product.space = theSize * (1 + aditionalPercentageSpace[theProductType])
	if product.space > SHIPPING_LIMIT {
		return product, errors.New(theName + " occupies too much space, sorry")
	}

	return product, nil
}

func (product Product) getSpaceRequired() float64 {
	return product.space
}

type ShippingFee struct {
	products     []Product
	numShippings int
}

type ShippingBuilder struct {
	shippingFee ShippingFee
	spaceNeeded float64
}

func (shippingBuilder ShippingBuilder) addProduct(product Product) ShippingBuilder {

	shippingBuilder.shippingFee.products = append(shippingBuilder.shippingFee.products, product)
	productSpace := product.getSpaceRequired()
	aux := shippingBuilder.spaceNeeded + productSpace
	if aux > SHIPPING_LIMIT {
		shippingBuilder.shippingFee.numShippings++
		shippingBuilder.spaceNeeded = productSpace

	} else {
		shippingBuilder.spaceNeeded = aux
	}
	return shippingBuilder
}

func (shippingBuilder ShippingBuilder) build() ShippingFee {

	return shippingBuilder.shippingFee
}

func main() {
	builder := ShippingBuilder{shippingFee: ShippingFee{numShippings: 1}}

	product, err := newProduct("tv", FRAGILE, 1000)
	if err != nil {
		fmt.Println(err)
	} else {
		builder = builder.addProduct(product)
	}
	product, err = newProduct("socks", SMALL, 200)
	if err != nil {
		fmt.Println(err)
	} else {
		builder = builder.addProduct(product)
	}
	product, err = newProduct("car", BIG, 9000000)
	if err != nil {
		fmt.Println(err)
	} else {
		builder = builder.addProduct(product)
	}
	product, err = newProduct("bike", MEDIUM, 1050)
	if err != nil {
		fmt.Println(err)
	} else {
		builder = builder.addProduct(product)
	}

	shipping := builder.build()

	fmt.Println(shipping.numShippings, "shippments for: ", shipping.products)

}
