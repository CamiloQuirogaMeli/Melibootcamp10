package main

import "fmt"

func main() {
	fmt.Print("Ingrese una palabra: ")
	var palabra string
	fmt.Scanln(&palabra)

	fmt.Println("La palabra tiene ", len(palabra), " letras")

	for i := 0; i < len(palabra); i++ {
		if i != len(palabra)-1 {
			fmt.Print("Caracter ", i+1, ": ", string(palabra[i]), " ")
		} else {
			fmt.Println(" Caracter ", i+1, ": ", string(palabra[i]))
		}
	}
}
