package main

import "fmt"

func main() {
	var price float64 = 12345.0
	var discount int64 = 10
	fmt.Printf("%.2f\n", price-price*float64(discount)/100)
}
