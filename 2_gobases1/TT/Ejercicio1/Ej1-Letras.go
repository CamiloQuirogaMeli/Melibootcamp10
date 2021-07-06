package main

import "fmt"

func main() {
	var word string = "diccionario"

	fmt.Println(len(word))

	fmt.Println("Letras de la palabra: ")

	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]))
	}
}
