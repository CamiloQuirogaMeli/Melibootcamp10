package main

import "fmt"

func main() {
	var word = "perro"
	var charCount = len([]rune(word))
	fmt.Println("Cantidad de letras: ", charCount)
	fmt.Println("Deletreando palabra: ", word)
	for _, l := range word {
		fmt.Println(string(l))
	}
}
