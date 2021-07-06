package main

import "fmt"

func main() {
	word := "palabras"
	fmt.Println("Cantidad de letras de la palabra: ", len(word))

	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]))
	}
}
