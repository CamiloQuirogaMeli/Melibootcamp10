package main

import "fmt"

func main() {
	fmt.Println("Ingrese una palabra para analizar: ")
	var palabra string
	fmt.Scanln(&palabra)
	fmt.Println("La palabra '", palabra, "' tiene ", len(palabra), " letras")
	fmt.Println("Las letras son:")
	for i := 0; i < len(palabra); i++ {
		fmt.Println(palabra[i : i+1])
	}
}
