package main

import "fmt"

func main() {
	var precio float32 = 24
	var descuento float32 = 0.15

	total := precio * (1 - descuento)

	fmt.Println(total)
}
