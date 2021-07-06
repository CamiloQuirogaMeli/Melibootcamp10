// Ejercicio 1 - Letras de una palabra

// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada
// una de las letras por separado para deletrearla, para ello es necesario crear una aplicación
// tener una variable con la palabra e imprimir la cantidad de letras que tiene la misma, luego
// imprimí cada una de las letras
package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Printf("Escribe una palabra: ")
	fmt.Scanf("%s\n", &word)
	fmt.Println(word, "(", len(word), " letras): ")
	for i, letter := range word {
		fmt.Println(i, "- ", string(letter))
	}
	fmt.Println()
}
