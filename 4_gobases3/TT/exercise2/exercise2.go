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

func createProduct(name *string, price *float64) Product {
	return Product{
		name:  *name,
		price: *price,
	}
}

func addProduct(user *User, product *Product, quantity *int) {
	(*product).quantity = *quantity
	(*user).products = append((*user).products, *product)
}

func deleteProduct(user *User) {
	(*user).products = []Product{}
}

func main() {
	var (
		name         = "Carlos"
		lastName     = "Infante"
		email        = "carlos.infante@mercadolibre.com.co"
		productName  = "Milanesa"
		productPrice = 5000.0
		quantity     = 10
	)
	user := User{
		name:     name,
		lastName: lastName,
		email:    email,
	}

	newProduct := createProduct(&productName, &productPrice)
	addProduct(&user, &newProduct, &quantity)
	fmt.Println(user)
	deleteProduct(&user)
	fmt.Println(user)
}
