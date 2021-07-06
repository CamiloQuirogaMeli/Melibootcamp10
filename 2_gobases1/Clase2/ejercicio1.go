package main

import (
	"fmt"
	"strings"
)

func main() {

	/*
		La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada
		una de las letras por separado para deletrearla, para ello es necesario crear una aplicación
		tener una variable con la palabra e imprimir la cantidad de letras que tiene la misma, luego
		imprimí cada una de las letras.
	*/

	var palabra string

	fmt.Print("Digita una palabra: ")
	fmt.Scanln(&palabra)

	fmt.Println("La palabra tiene ", len(palabra), "letras")
	fmt.Println("Las letras separadas son: ")

	letras := strings.Split(palabra, "")

	for num, letra := range letras {
		fmt.Println("Letra ", num+1, ":", letra)
	}

}
