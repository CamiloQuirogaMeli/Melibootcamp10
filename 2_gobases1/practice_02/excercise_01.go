package main

import "fmt"

func main() {

	var word string
	fmt.Println("Please, input one word")
	fmt.Scanln(&word)

	fmt.Printf("Word length: [%d]\nCharacters \n", len(word))
	for char := range word {
		fmt.Println(string(word[char]))
	}
}
