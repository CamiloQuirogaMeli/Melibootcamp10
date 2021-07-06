package main

import "fmt"

func main() {
	fmt.Println("Escribe una palabra: ")
	var word string
	fmt.Scanln(&word)
	fmt.Printf("La palabra %s tiene %d letras\n", word, len(word))
	for _, char := range word {
		fmt.Printf("%c\n", char)
	}
}
