package main

import "fmt"

type users struct {
	name string
	lastName string
	mail string
	products []product
}

type product struct {
	name string
	price float64
	amount int
}

func main() {
	product := newProduct("toalla", 14.6, 5)
	fmt.Println("el producto es:", *product)
	user := &users{
		name: "pepe",
		mail: "pepe@ml.com",
		lastName: "toto",
	}
	addProduct(user, product, product.amount)
	fmt.Println(*user)
	product2 := newProduct("galletitas", 99.9, 99)
	addProduct(user, product2, product2.amount)
	fmt.Println(*user)
	clearProducts(user)
	fmt.Println(*user)
}

func newProduct(name string, price float64, amount int) *product {
	return &product{
		name: name,
		price: price,
		amount: amount,
	}
}

func addProduct(u *users, p *product, amount int) {
	u.products = append(u.products, *p)
}

func clearProducts(u *users){
	u.products = nil // [:0]
}