package main

import "fmt"

func main() {

	fmt.Println("Type a word:")
	var word string

	fmt.Scanln(&word)

	fmt.Println("Length: ", len(word))
	fmt.Println("Letters:")

	for i := range word {
		fmt.Println("\t", string(word[i]))
	}

}
