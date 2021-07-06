package main

import (
	"fmt"
	"strings"
)

func cantidadLetras(word string) {

	var letters []string = strings.Split(word, "")

	for i := 0; i < len(word); i++ {
		fmt.Println(letters[i])
	}

	fmt.Printf("Cantidad de letras: %d\n", len(word))
}
