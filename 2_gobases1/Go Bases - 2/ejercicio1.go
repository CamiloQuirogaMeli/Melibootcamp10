package main

import (
	"fmt"
)

/*
	Ejercicio 1 - Letras de una palabra
		La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras
		por separado para deletrearla, para ello es necesario crear una aplicación tener una variable con la palabra
		e imprimir la cantidad de letras que tiene la misma, luego imprimí cada una de las letras.
*/

func main() {
	palabra := "prueba"
	letras := []byte(palabra)
	fmt.Printf("La cantidad de letras para %q son %d\n", palabra, len(palabra))
	for i, letra := range letras {
		fmt.Printf("La letra %d es igual a: %c\n", i+1, letra)
	}
}
