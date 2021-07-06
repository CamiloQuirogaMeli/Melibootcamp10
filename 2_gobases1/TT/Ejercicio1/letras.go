package main

import "fmt"

func main() {

	var palabra string = "hola"

	fmt.Println(len(palabra))

	for _, elem := range palabra {
		fmt.Println(string(elem))
	}
}
