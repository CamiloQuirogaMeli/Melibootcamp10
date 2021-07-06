package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string = "almuerzo"
	fmt.Println("Cantidad de letras: ", len(word))

	letters := strings.Split(word, "")
	fmt.Println(letters)
	for i, letter := range letters {
		fmt.Println(i, letter)
	}
}
