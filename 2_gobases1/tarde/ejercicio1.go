package main

import "fmt"

func main() {
	var word string = "Mercadolibre"

	fmt.Println("La palabra", word, "tiene ", len(word), "letras")

	fmt.Println("A continuación cada una de sus letras: ")
	for _, letter := range word {
		fmt.Printf("%c \n", letter)
	}

}
