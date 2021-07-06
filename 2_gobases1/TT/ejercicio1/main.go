package main

import "fmt"

func main() {

	letras("Hola")

}

func letras(palabra string) {

	var cant int = len(palabra)

	fmt.Println("La palabra "+palabra+
		" tiene ", cant, " letras.")

	for _, r := range palabra {
		c := string(r)
		fmt.Println(c)
	}

}
