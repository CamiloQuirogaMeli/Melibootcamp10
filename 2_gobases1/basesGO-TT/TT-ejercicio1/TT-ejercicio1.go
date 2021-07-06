package main

import "fmt"

/*
Ejercicio 1 - Letras de una palabra

La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada
una de las letras por separado para deletrearla, para ello es necesario crear una aplicación
tener una variable con la palabra e imprimir la cantidad de letras que tiene la misma, luego
imprimí cada una de las letras.
*/

func tenerUnaVariable(palabra string) {

	/*Imprimo el tamaño de la palabra*/
	fmt.Println("Tamaño de la palabra: ", len(palabra))

	/*Imprimo cada una de las letras*/
	fmt.Println("Letras que la componenen: ")

	for i := 0; i < len(palabra); i++ {
		/*fmt.Printf("%q", palabra[i])*/
		fmt.Println(string(palabra[i]))
	}

	/* PROBANDO OTRA SOLUCION

	for j, pal := range palabra {
		fmt.Printf("%q", pal)
		j++
	}

	*/

}

func main() {

	palabra := "Hola"
	fmt.Println("Palabra: ", palabra)
	tenerUnaVariable(palabra)

}
