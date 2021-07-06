package main

import "fmt"

func main() {

	//
	// Ejercicio 1 - Letras de una palabra
	// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada
	// una de las letras por separado para deletrearla, para ello es necesario crear una aplicación
	// tener una variable con la palabra e imprimir la cantidad de letras que tiene la misma, luego
	// imprimí cada una de las letras.
	//

	palabra := "Hola"
	tam := len(palabra)
	fmt.Println("Palabra: ", palabra)
	fmt.Println("Total de letras: ", tam)
	for i := 0; i < tam; i++ {
		fmt.Println("Letra: ", string(palabra[i]))
	}
}
