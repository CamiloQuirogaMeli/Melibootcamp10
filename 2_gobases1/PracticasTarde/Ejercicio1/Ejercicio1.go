package main

import "fmt"

func main() {

	/*
		Ejercicio 1 - Letras de una palabra
		La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada
		una de las letras por separado para deletrearla, para ello es necesario crear una aplicación
		tener una variable con la palabra e imprimir la cantidad de letras que tiene la misma, luego
		imprimí cada una de las letras.
	*/
	fmt.Println("Ingrese una palabra")
	var word string
	fmt.Scanln(&word)

	fmt.Println("La palabra \"", word, "\" tiene ", len(word), " caracteres.")
	//Forma 1
	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]))
	}
	
	//Forma 2
	for _, letter := range word {
		fmt.Println(string(letter))
	}
}
