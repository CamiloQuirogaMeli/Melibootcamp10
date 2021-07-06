package main

import "fmt"

func main() {
	var word string = "mercadoLibre"

	fmt.Printf("El numero de letras de la palabra \"%s\" es %d: \n", word, len(word))

	fmt.Println("Se deletrea: ")
	for i := 0; i < len(word); i++ {
		fmt.Println(word[i : i+1])
	}
}
