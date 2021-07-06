package main

import "fmt"

func main() {

	var completeWord string = "Otorrinolaringologia"

	fmt.Printf("La palabra seleccionada es %v y tiene %v letras \n", completeWord, len(completeWord))

	for key, element := range completeWord {
		fmt.Printf("En la posición %v está la letra %c \n", key, element)
	}

}
