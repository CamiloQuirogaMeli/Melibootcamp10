package main

import "fmt"

func main() {
	var word string = "murcielago"
	fmt.Printf("%d\n", len(word))
	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]))
	}
}
