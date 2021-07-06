package main

import "fmt"

func main() {
	var word string = "RANDOM"
	var lenWord int = len(word)
	fmt.Println("len: ", lenWord)

	for _, letter := range word {
		fmt.Println("letter", string(letter))
	}

}
