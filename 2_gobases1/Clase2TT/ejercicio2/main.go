package main

import "fmt"

func main() {

	var price float64 = 4000
	var discount float64 = 50

	var totalPrice float64 = price - (price * discount / 100)

	fmt.Println(totalPrice)
}
