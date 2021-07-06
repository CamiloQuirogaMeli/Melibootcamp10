package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	mail      string
	products  []product
}

type product struct {
	name     string
	price    float64
	quantity int
}

func newProduct(name *string, price *float64) product {
	return product{name: *name, price: *price}
}

func addProduct(user *user, product *product, quantity *int) {

	product.quantity = *quantity

	user.products = append(user.products, *product)
}

func deleteProduct(user *user) {

	user.products = []product{}

}

func main() {

	nameProd := "prod1"
	priceProd := 123.00

	prod2 := newProduct(&nameProd, &priceProd)

	user := user{"John", "Doe", "john.doe@asd.com", []product{}}

	quantityProd := 10

	addProduct(&user, &prod2, &quantityProd)

	fmt.Println(user)

	deleteProduct(&user)

	fmt.Println(user)

}
