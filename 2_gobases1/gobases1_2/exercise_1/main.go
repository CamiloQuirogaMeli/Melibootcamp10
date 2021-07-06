package main

import (
	"fmt"
	"strings"
)

var (
	word string
)

func main() {

	fmt.Printf("Enter a word: ")
	fmt.Scanln(&word)

	// count the letters of the word
	len_word := len(word)
	fmt.Println("\nNumber of letters in the word: ", len_word)

	// Slice word into letters:

	// method 1
	for i, letter := range word {
		fmt.Printf("\nl_%d : %c \n", i, letter)
	}

	// method 2:

	str := strings.Split(word, "")

	fmt.Printf("\n%q", str)
}
