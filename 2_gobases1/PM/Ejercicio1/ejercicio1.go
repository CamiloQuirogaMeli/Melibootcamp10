package main

import "fmt"

func main() {
	var palabra string = "Aquelarre"
	var cantidad int = len(palabra)
	fmt.Println(cantidad)
	for i, r := range palabra {
		fmt.Printf("%d %c\n", i+1, r)
	}
}
