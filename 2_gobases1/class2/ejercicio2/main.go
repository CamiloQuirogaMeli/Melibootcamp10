package main

import "fmt"

func main() {
	discountPrice(585, 13)
}

func discountPrice(price, discount int) {
	totalPrice := float64(discount * price)
	totalPrice /= 100
	totalPrice = float64(price) - totalPrice
	fmt.Println(totalPrice)
}
