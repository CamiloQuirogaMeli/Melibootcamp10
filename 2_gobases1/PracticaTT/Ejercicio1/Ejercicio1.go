// Ejercicio 1 - Letras de una palabra

// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada
// una de las letras por separado para deletrearla, para ello es necesario crear una aplicación
// tener una variable con la palabra e imprimir la cantidad de letras que tiene la misma, luego
// imprimí cada una de las letras.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var word string = "Estudios"
	var wordLength int
	var valueSpace bool = true
	scanner := bufio.NewScanner(os.Stdin)
	s
	for valueSpace == true {
		fmt.Printf("Ingresar una palabra: ")
		scanner.Scan()
		word = scanner.Text()
		fmt.Println("Word:", word)
		valueSpace = strings.Contains(word, " ")
	}

	wordLength = len(word)
	fmt.Println("La palabra que escribiste tiene :", wordLength, "letras")
	fmt.Print("La palabra separada:\n")
	for i := 0; i < wordLength; i++ {
		fmt.Printf("Letra %d: %c\n", i+1, word[i])
	}
}
