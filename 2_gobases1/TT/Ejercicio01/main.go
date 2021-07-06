package main

import "fmt"

func main() {
	var word string
	fmt.Print("Ingrese una palabra: ")
	fmt.Scanf("%s", &word)

	fmt.Println("Cantidad de letras:", len(word))

	for i, char := range word {
		fmt.Println("Letra:", string(char), "- Posicion:", i+1)
	}

	/* for i := 0; i < len(word); i++ {
		fmt.Println("Letra:", string(word[i]), "Posicion:", i+1)
	}*/
}
