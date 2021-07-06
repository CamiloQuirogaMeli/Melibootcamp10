package main

import (
	"fmt"
)

type Product struct {
	name     string
	price    float64
	quantity int64
}

type User struct {
	name     string
	lastName string
	email    string
	products []Product
}

func newProduct(name *string, price *float64) Product {
	return Product{(*name), (*price), 0}
}

func addProduct(u *User, p *Product, q *int64) {
	(*p).quantity = (*q)
	(*u).products = append((*u).products, (*p))
}

func deleteProduct(u *User) {
	(*u).products = []Product{}
}

func main() {
	usr := User{"will", "Garcia", "t@t.com", []Product{}}
	fmt.Println(usr)
	productName, productPrice := "product1", 1235.00
	p := newProduct(&productName, &productPrice)
	var quantity int64 = 3
	addProduct(&usr, &p, &quantity)
	fmt.Println(usr)
	deleteProduct(&usr)
	fmt.Println(usr)
}
