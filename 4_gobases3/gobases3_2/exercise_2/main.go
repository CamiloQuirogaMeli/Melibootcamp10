package main

import (
	"fmt"
)

var (
	log          string
	name         string
	price        float64
	stock        int
	act          string
	products     []Product
	shoppingCart []Product
	user_pointer *User
)

func main() {

	for {
		fmt.Println("you want to log in as a user or as an employee? ")
		fmt.Scanln(&log)

		if log == "employee" {

			i := 1
			for i == 1 {
				fmt.Print("You want to enter a new product? (1 yes - 2 no): ")
				fmt.Scanln(&i)

				if i == 1 {
					fmt.Print("Enter the name: ")
					fmt.Scanln(&name)
					fmt.Print("Enter the price: ")
					fmt.Scanln(&price)
					fmt.Print("Enter the stock: ")
					fmt.Scanln(&stock)

					products = append(products, newProduct(name, stock, price))
					fmt.Println(products)
				} else {
					break
				}
			}
		} else if log == "user" {

			for {

				user := User{"Joseph", "Caicedo", "joseph.saenz@mercadolibre.com.co", shoppingCart}
				user_pointer = &user

				fmt.Println("\nWhat action do you want to take?")
				fmt.Println("1. Add product to cart")
				fmt.Println("2. Delete all products from the cart")
				fmt.Println("Enter any character or number to exit\n")
				fmt.Scanln(&act)

				switch act {
				case "1":
					addProduct(user_pointer, products[0], 1)
					fmt.Println(user)
				case "2":
					delateProducts(user_pointer)
					fmt.Println(user)
				default:
					break
				}

			}
		} else {
			break
		}
	}
}
