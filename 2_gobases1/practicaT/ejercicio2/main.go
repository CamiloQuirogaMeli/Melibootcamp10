package main

import "fmt"

func main() {

	var price int = 1000
	var avgDiscount int = 10
	var sale int = price * (100 - avgDiscount) / 100

	fmt.Println("sale", sale)
}
