package main

import "fmt"

func main() {
	palabra := "Agricultura"
	palabraArray := []rune(palabra)
	fmt.Printf("El tamaño de la palabra es %d \n", len(palabraArray))

	for i := 0; i < len(palabraArray); i++ {
		fmt.Println(string(palabraArray[i]))
	}

}
