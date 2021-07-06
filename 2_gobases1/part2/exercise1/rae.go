package main

import (
	"fmt"
)

func main() {
	var word string = "Paralelogramo"
	fmt.Println("Word length: ", len(word))
	fmt.Println("Spelling of the word: ")
	for i, letter := range word {
		fmt.Printf("%d: %c\n", i, letter)
	}
}

