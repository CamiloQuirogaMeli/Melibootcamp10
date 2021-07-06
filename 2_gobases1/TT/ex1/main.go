package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Print("Ingrese la palabra: ")
	fmt.Scanln(&word)
	fmt.Println("La palabra", word, "tiene", len(word), "letras")
	fmt.Println("Deletreada:")
	for i := 0; i < (len(word) - 1); i++ {
		fmt.Print(string(word[i]), " ")
	}
	fmt.Println(string(word[len(word)-1]))
}
