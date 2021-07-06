package main

import (
	"fmt"
)

type Product struct {
	Name   string
	Price  float64
	Amount int
}

type User struct {
	Name     string
	Surname  string
	Email    string
	Products []Product
}

func newProduct(name string, price float64) Product {
	return Product{Name: name, Price: price}
}

func addProduct(user *User, product *Product, amount int) {
	product.Amount = amount
	user.Products = append(user.Products, *product)
}

func deleteProducts(user *User) {
	user.Products = nil
}

func main() {
	user := User{Name: "Juan", Surname: "Perez", Email: "ju@gmail.com"}
	for i := 0; i < 5; i++ {
		p := newProduct("producto "+fmt.Sprint(i), float64(i)+10)
		addProduct(&user, &p, 5)
	}
	fmt.Println(user)
	deleteProducts(&user)
	fmt.Println(user)
}
