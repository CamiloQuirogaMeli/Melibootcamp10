package main

import "fmt"

func main() {
	var palabra string = "palabra"
	fmt.Println("Cantidad de palabras: ", len(palabra))
	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%c\n", palabra[i])
	}
}
