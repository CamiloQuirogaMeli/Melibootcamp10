package main

import "fmt"

var (
	price       float32
	discount    float32
	final_price float32
)

func main() {

	fmt.Print("Enter product price: ")
	fmt.Scanln(&price)

	fmt.Print("Enter product discount: ")
	fmt.Scanln(&discount)

	final_price = price - price*discount/100
	fmt.Printf("\nThe price of the product with the discount applied is: %.2f", final_price)

}
