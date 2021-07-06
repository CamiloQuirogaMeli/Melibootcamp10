package main

import "fmt"

func main() {
	var a *int
	b := 10
	a = &b
	fmt.Println(&b) // dirección de memoria de b (1008)
	fmt.Println(a)  // guarda la dirección de memoria de b (1008)
	fmt.Println(b)  // valor de b (10)
	fmt.Println(*a) // valor a donde apunta a (10)
}
