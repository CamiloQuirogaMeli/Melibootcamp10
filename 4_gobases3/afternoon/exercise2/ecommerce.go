package main

import "fmt"

type User struct {
	name     string
	age      int
	email    string
	products []*Product
}

func (user *User) String() string {
	return fmt.Sprint("Name: ", user.name, "\nProducts: ", user.products)
}

type Product struct {
	name   string
	price  float64
	amount int
}

func (product *Product) String() string {
	return product.name
}

func main() {

	user := User{name: "Andres", age: 24, email: "afmoyar"}
	p1 := newProduct("poni malta", 1500)
	p2 := newProduct("choquis", 1000)
	p3 := newProduct("chiclex", 200)
	p4 := newProduct("chocorramo", 2000)
	p5 := newProduct("bon yout", 1500)

	addUserProduct(&user, p1, 2)
	addUserProduct(&user, p2, 2)
	addUserProduct(&user, p3, 2)
	addUserProduct(&user, p4, 2)
	addUserProduct(&user, p5, 2)

	fmt.Println(user.String())

	resetUserProducts(&user)
	fmt.Println(user.String())
}

func newProduct(theName string, thePrice float64) *Product {
	return &Product{name: theName, price: thePrice}
}

func addUserProduct(user *User, product *Product, theAmount int) {
	product.amount = theAmount
	user.products = append(user.products, product)
}

func resetUserProducts(user *User) {
	user.products = []*Product{}
}
