package main

import "fmt"

func main() {
	var word string = "palabra"
	fmt.Printf("Cantidad de letras de la palabra: %d", len(word))
	for i, letter := range word {
		fmt.Printf("\n%d- %c ", i+1, letter)
	}
}
