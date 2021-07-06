package main

import (
	"fmt"
	"strings"
)

func main() {
	var palabra string

	fmt.Println("Ingrese palabra a deletrear: ")

	fmt.Scanln(&palabra)
	palabra = strings.ToLower(palabra)

	fmt.Println("Cantidad palabra ingresada tiene ", len(palabra), " letras")
	for _, letra := range palabra {
		fmt.Println(string(letra))
	}

}
