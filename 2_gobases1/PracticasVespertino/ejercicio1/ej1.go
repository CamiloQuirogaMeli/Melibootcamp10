package main

import "fmt"

var word = "HolaSoyUnaPalabra"

func main() {

	fmt.Printf("\nLa cantidad de letras de la palabra %q es: %d\n\n", word, len(word))

	for char := range word {
		fmt.Println(string(word[char]))
	}
}

/*
	 La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada
	una de las letras por separado para deletrearla, para ello es necesario crear una aplicación
	tener una variable con la palabra e imprimir la cantidad de letras que tiene la misma, luego
	imprimí cada una de las letras.
*/
