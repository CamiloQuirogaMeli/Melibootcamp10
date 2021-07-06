package main

import "fmt"

func main() {

	myWord := "Externocleidomastoideo"
	size := len(myWord)
	fmt.Println("The word has", size, "words\nNow let's spell the word, shall we?")

	for index, rune := range myWord {
		if index == size-1 {
			fmt.Print(string(rune))
		} else {
			fmt.Print(string(rune), ",")
		}

	}
}
