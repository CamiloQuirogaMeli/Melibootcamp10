package main

import (
	"fmt"
)

func spellOutWord(word string) string {
	var spelledWord string = ""
	for _, c := range word {
		spelledWord += string(c) + " "
	}
	return spelledWord
}

func main() {

	var word string = "paralepipedo"

	fmt.Printf("La longitud de la palabra %s es %d\n", word, len(word))
	fmt.Printf("La palabra %s se deletrea %s\n", word, spellOutWord(word))
}