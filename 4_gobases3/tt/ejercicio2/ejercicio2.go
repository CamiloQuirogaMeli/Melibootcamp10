package main

import "fmt"

type Product struct {
	name  string
	price float64
	total int
}

type User struct {
	name     string
	surname  string
	email    string
	products []*Product
}

func newProduct(n string, p float64) *Product {
	return &Product{
		name:  n,
		price: p,
	}
}

func addProduct(user *User, p *Product, t int) {
	ok := 0
	for _, value := range user.products {
		if value.name == p.name && value.price == p.price {
			value.total += t
			ok = 1
		}
	}
	if ok == 0 {
		p.total = t
		user.products = append(user.products, p)
	}
}

func deleteProduct(user *User) {
	user.products = []*Product{}
}

func main() {
	user := &User{
		name:    "Karen",
		surname: "Cuadrado",
		email:   "cuadrado.karen@gmail.com",
	}

	p1 := newProduct("producto 1", 123.45)
	p2 := newProduct("producto 2", 432.05)
	p3 := newProduct("producto 3", 856.05)
	addProduct(user, p1, 10)
	addProduct(user, p2, 3)
	addProduct(user, p1, 5)
	addProduct(user, p3, 5)
	fmt.Println(user.name, user.surname, user.email)
	for _, value := range user.products {
		fmt.Println(*value)
	}
}
