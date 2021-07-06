package main

import (
	"fmt"
)

func ej1LP() {
	fmt.Print("Ingrese una palabra: ")
	var wordToCount string
	fmt.Scanln(&wordToCount) //estoy haciendo que ingrese por consola una palabra
	//wordToCount := "Diego"

	for i, letter := range wordToCount {
		fmt.Printf("%d %c\n", i, letter)

	}
	fmt.Println("La cantidad de letras que tiene la palabra es ", len(wordToCount))

	//arreglo que no hizo falta
	/*for marc < len(wordToCount) {


		marcCount++

		marc += len(wordToCount)

	}*/

}
