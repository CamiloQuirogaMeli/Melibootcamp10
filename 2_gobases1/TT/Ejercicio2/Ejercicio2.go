package main

import "fmt"

func main() {
	var price, discount float64 = 134, 25
	a := price * (1 - (discount / 100))
	fmt.Println("Precio final:", a)
}
