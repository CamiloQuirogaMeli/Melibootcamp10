package main

import "fmt"

func main() {
	var palabra string = "ejercicio1"

	println("Esta palabra tiene: ", len(palabra), "letras")

	for i := 0; i < len(palabra); i++ {
		fmt.Println(string(palabra[i]))
	}
}
