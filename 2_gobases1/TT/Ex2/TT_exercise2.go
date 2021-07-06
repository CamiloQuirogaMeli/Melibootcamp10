package main

import "fmt"

func main() {

	var price, discount, total float32

	fmt.Println("Enter the price:")
	fmt.Scanln(&price)

	fmt.Println("Enter the discount(%):")
	fmt.Scanln(&discount)

	total = price - (price * (discount / 100))

	fmt.Println("Final Amount: ", total)

}
