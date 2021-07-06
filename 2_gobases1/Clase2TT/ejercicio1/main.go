package main

import "fmt"

func main() {

	var word string = "word"

	fmt.Println(len(word))

	for _, letter := range word {
		fmt.Printf("%c \n", letter)
	}
}
