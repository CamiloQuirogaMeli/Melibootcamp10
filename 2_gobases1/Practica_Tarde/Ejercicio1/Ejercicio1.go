package main

import "fmt"

func main() {
	var palabra string = "Argentina"
	var cantidad = len(palabra)

	fmt.Println("Cantidad de palabras: ", cantidad)
	for i := 0; i < cantidad; i++ {
		fmt.Printf("%c\n", palabra[i])
	}

}
