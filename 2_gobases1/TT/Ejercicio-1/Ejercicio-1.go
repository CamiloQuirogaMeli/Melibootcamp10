package main

import "fmt"

func printWord(Word string) {
	length := len(Word)
	fmt.Println("Word length of "+Word+" is: ", length)
	for i, letra := range Word {
		fmt.Println("Posicion: ", i+1, ", Letra: "+string(letra))
	}
}

func main() {
	fmt.Println("Ejercicio 1")
	printWord("esternocleidomastoideo")
}
