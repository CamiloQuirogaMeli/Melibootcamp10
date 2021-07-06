package main

import "fmt"

func main() {

	var word string = "palabraDePrueba"
	var lettersQuantity int = len(word)

	fmt.Printf("La palabra tiene %d letras", lettersQuantity)

	for _, lettersQuantity := range word {
		fmt.Printf("%c \n", lettersQuantity)
	}

}
