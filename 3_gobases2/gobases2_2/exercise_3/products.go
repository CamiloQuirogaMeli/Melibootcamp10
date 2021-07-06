package main

import (
	er "errors"
	f "fmt"
)

type Product struct {
	product_name  string
	product_type  string
	product_price float64
}

type Shop struct {
	shop_address     string
	product          Product
	delivery_address string
}

type Logitech struct {
	shop Shop
}

func (l Logitech) calculate_price() (float64, error) {

	var additional_cost float64

	if l.shop.product.product_type == "small" {
		additional_cost = 0
	} else if l.shop.product.product_type == "medium" {
		additional_cost = 0.03 * l.shop.product.product_price
	} else if l.shop.product.product_type == "big" {
		additional_cost = 0.06*l.shop.product.product_price + 2500
	} else {
		return 0, er.New("This type of product does not exist")
	}

	return additional_cost + l.shop.product.product_price, nil
}

func (l Logitech) send_product() string {
	return l.shop.delivery_address
}

type HyperX struct {
	shop Shop
}

func (h HyperX) calculate_price() (float64, error) {

	var additional_cost float64

	if h.shop.product.product_type == "small" {
		additional_cost = 0
	} else if h.shop.product.product_type == "medium" {
		additional_cost = 0.03 * h.shop.product.product_price
	} else if h.shop.product.product_type == "big" {
		additional_cost = 0.06*h.shop.product.product_price + 2500
	} else {
		return 0, er.New("This type of product does not exist")
	}

	return additional_cost + h.shop.product.product_price, nil
}

func (h HyperX) send_product() string {
	return h.shop.delivery_address
}

type Ecommerce interface {
	calculate_price() (float64, error)
	send_product() string
}

const (
	logitech = "LOGITECH"
	hyperX   = "HYPERX"
)

func newShop(shop_name string, shop_address string, product Product, delivery_address string) Ecommerce {

	shop := Shop{shop_address, product, delivery_address}

	switch shop_name {
	case logitech:
		return Logitech{shop}
	case hyperX:
		return HyperX{shop}
	}
	return nil
}

func newProduct() Product {

	product := Product{}

	f.Print("Enter the product name: ")
	f.Scanln(&product.product_name)

	f.Print("Enter the product type: ")
	f.Scanln(&product.product_type)

	f.Print("Enter the product price: ")
	f.Scanln(&product.product_price)

	return product
}

func main() {
	product := newProduct()
	l := newShop(logitech, "Cal 20B N45B - 45", product, "Cr 35 60A N65")
	total_price1, err1 := l.calculate_price()
	handleErrors(total_price1, err1)
	f.Printf("The product was shipped to %s \n\n", l.send_product())

	product2 := newProduct()
	h := newShop(hyperX, "Tv 3 N6 45", product2, "Cl 37 34B - 35")
	total_price2, err2 := h.calculate_price()
	handleErrors(total_price2, err2)
	f.Printf("The product was shipped to %s \n\n", h.send_product())
}

func handleErrors(total_price float64, err error) {
	if err == nil {
		f.Println("The total price is: ", total_price)
	} else {
		f.Println(err)
	}
}
