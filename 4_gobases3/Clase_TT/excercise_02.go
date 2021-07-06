package main

import (
	"fmt"
)

func main() {
	var pointerUser *UserA
	user := UserA{Name: "Sebastian", LastName: "Posadas", Email: "sebastian.posadas@mercadolibre.com", Products: []Product{}}
	pointerUser = &user

	newProduct := newProduct("Mouse", 2500.00)
	fmt.Println(newProduct)

	addProduct(pointerUser, newProduct, 5)

	fmt.Println(user)

	removeProducts(pointerUser)

	fmt.Println(user)
}

type UserA struct {
	Name     string
	LastName string
	Email    string
	Products []Product
}

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func newProduct(name string, price float64) *Product {
	var pointerProduct *Product
	newProduct := Product{Name: name, Price: price}
	pointerProduct = &newProduct
	return pointerProduct
}

func addProduct(u *UserA, p *Product, quantity int) {
	(*p).Quantity = quantity
	(*u).Products = append((*u).Products, *p)
}

func removeProducts(u *UserA) {
	(*u).Products = nil
}
