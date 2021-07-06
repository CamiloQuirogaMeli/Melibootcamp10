package main

import "fmt"

type User struct {
	name     string
	lastName string
	email    string
	products []Product
}

type Product struct {
	name     string
	price    float64
	quantity int
}

func newProduct(name string, price float64) Product {

	return Product{name, price, 1}
}

func addProduct(user *User, product *Product, quantity int) {

	for i := 0; i <= quantity; i++ {
		user.products = append(user.products, *product)
	}
}

func deleteProducts(user *User) {
	user.products = make([]Product, 0)
}

func main() {

	product1 := newProduct("doritos", 200)
	product2 := newProduct("monster", 150)

	user := User{"Gianluca", "Ciccarelli", "email@gmail.com", make([]Product, 0)}

	addProduct(&user, &product1, 2)
	addProduct(&user, &product2, 4)

	fmt.Println(user)

	deleteProducts(&user)

	fmt.Println(user)
}
