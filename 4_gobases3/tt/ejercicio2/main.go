package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Product struct {
	name   string
	price  float64
	amount int
}

type User struct {
	name     string
	email    string
	products []Product
}

func NewProduct(name string, price float64) (*Product, error) {
	if name == "" {
		return nil, errors.New("no name provided")
	}
	if price < 0 {
		return nil, errors.New("invalid price")
	}

	p := &Product{
		name:   name,
		price:  price,
		amount: 0,
	}
	return p, nil
}

func (u *User) AddProduct(p Product, amount int) {
	if amount <= 0 {
		amount = 1
	}
	p.amount = amount
	u.products = append(u.products, p)
}

func (u *User) RemoveProducts() {
	u.products = nil
}

func (u *User) Info() {
	fmt.Printf("Name: \t %s\nEmail: \t %s\nProducts: \t %+v\n", u.name, u.email, u.products)
}

func main() {
	user := User{
		name:  "Marcos",
		email: "marcos.zabala@mercadolibre.com",
	}
	product, err := NewProduct("AMD Ryzen 5 3rd Gen", 30000.0)

	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	user.AddProduct(*product, 1)
	user.Info()
}
