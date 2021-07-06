package main

import (
	f "fmt"
)

type User struct {
	Name     string
	Lastname string
	Email    string
	Products []Product
}

type Product struct {
	Name     string
	Price    float32
	Quantity int
}

func newProduct(name string, price float32) Product {
	prod := Product{Name: name, Price: price}
	return prod
}

func (product *Product) addProduct(user *User, quantity int) {
	user.Products = append(user.Products, Product{Name: product.Name, Price: product.Price, Quantity: quantity})
}

func (product *Product) deleteProduct(user *User) {
	user.Products = []Product{}
}

func main() {
	user1 := User{Name: "Alejandro", Lastname: "Alamar", Email: "ale.jaam7@gmail.com"}
	user2 := User{Name: "Javier", Lastname: "Martinez", Email: "javier.martinez@gmail.com"}

	product := newProduct("Prod1", 1)
	product2 := newProduct("Prod2", 2)

	product.addProduct(&user1, 1)
	product2.addProduct(&user2, 2)

	f.Println("User 1:", user1.Name, user1.Lastname)
	f.Println("User 1 email:", user1.Email)
	f.Println("User 1 products:")
	for i, product := range user1.Products {
		f.Printf("%d.- %s (%d) - $%.2f\n", i+1, product.Name, product.Quantity, product.Price)
	}

	f.Println("User 2:", user2.Name, user2.Lastname)
	f.Println("User 2 email:", user2.Email)
	f.Println("User 2 products:")
	for i, product := range user2.Products {
		f.Printf("%d.- %s (%d) - $%.2f\n", i+1, product.Name, product.Quantity, product.Price)
	}

	product.deleteProduct(&user1)

	f.Println("After delete")
	f.Println("User 1:", user1.Name, user1.Lastname)
	f.Println("User 1 email:", user1.Email)
	f.Println("User 1 products:")
	for i, product := range user1.Products {
		f.Printf("%d.- %s (%d) - $%.2f\n", i+1, product.Name, product.Quantity, product.Price)
	}
}
