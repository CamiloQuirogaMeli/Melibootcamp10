package main

import "fmt"

type Product struct {
	name  string
	price float64
	stock int
}

type User struct {
	name     string
	surname  string
	email    string
	products []Product
}

func newProduct(name string, stock int, price float64) Product {

	product := Product{}

	product.name = name
	product.stock = stock
	product.price = price

	fmt.Println("New product added")

	return product
}

func addProduct(u *User, product Product, amount int) {
	if product.stock >= amount {
		u.products = append(u.products, product)
	} else {
		fmt.Println("No stock")
	}
}
func delateProducts(u *User) {
	u.products = u.products[:0]
}
