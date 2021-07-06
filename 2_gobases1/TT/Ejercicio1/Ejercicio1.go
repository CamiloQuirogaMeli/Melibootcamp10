package main

import "fmt"

func main() {
	var word = "Palabra"
	fmt.Println("Longitud:", len(word))

	for i, j := range word {
		fmt.Println(i, "=>", string(j))
	}
}
