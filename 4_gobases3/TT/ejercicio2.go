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
	return Product{name: name, price: price, quantity: 0}
}

func addProduct(user *User, product *Product, quantity int) {
	product.quantity = quantity
	user.products = append(user.products, *product)
}

func deleteProducts(user *User) {
	user.products = nil
}

func main() {
	var user = User{name: "Alex", lastName: "Castillo", email: "alejandro.castillo@mercadolibre.com.mx"}
	var user2 = User{name: "Alex2", lastName: "Castillo", email: "alejandro.castillo@mercadolibre.com.mx"}
	var product1 = newProduct("jabon", 20.0)
	var product2 = newProduct("detergente", 40.0)

	var product3 = newProduct("helado", 25.0)
	var product4 = newProduct("dulces", 15.0)

	addProduct(&user, &product1, 2)
	addProduct(&user, &product2, 3)
	fmt.Println(user)
	deleteProducts(&user)
	fmt.Println(user)

	addProduct(&user2, &product3, 1)
	addProduct(&user2, &product4, 20)
	fmt.Println(user2)
	deleteProducts(&user2)
	fmt.Println(user2)

}
