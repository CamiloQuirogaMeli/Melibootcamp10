package main

import "fmt"

func main() {
	originalPrice := 10000.0
	discount := 25.0
	newPrice := originalPrice * (1 - discount/100)

	fmt.Println("This shirt usually costs", originalPrice, "$")
	fmt.Println("But wait! We have a special discount of", discount, "%")
	fmt.Printf("So, if you want it now, it will cost only %.0f$", newPrice)
}
