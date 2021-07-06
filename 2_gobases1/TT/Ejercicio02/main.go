package main

import "fmt"

func main() {
	var (
		price      = 1000
		percentage = 30
		discount   = (price * percentage) / 100
	)

	fmt.Println("Precio final:", (price - discount))
}
