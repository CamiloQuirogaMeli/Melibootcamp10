package main

import (
	"fmt"
)

func main() {
	var palabra = "Palabra"
	fmt.Println(len(palabra))
	g := ""
	for i := 0; i < len(palabra); i++ {
		g = string(palabra[i])
		fmt.Println(g)
	}

}
