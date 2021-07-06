package main

import "fmt"

func main() {

	var price float64
	fmt.Println("Please, input the product price")
	fmt.Scanln(&price)

	var percentageDiscount float64
	fmt.Println("Please, input the percentage discount without percentage symbol")
	fmt.Scanln(&percentageDiscount)

	var discount float64 = price * (percentageDiscount / 100)

	fmt.Printf("The product price final is: [%2.f]\n", price-discount)
}
