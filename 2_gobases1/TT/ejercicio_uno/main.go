package main

import "fmt"

func main() {

	var slice = []string{"C", "A", "S", "A"}

	fmt.Printf("la palabra triene %d letras \n", len(slice))

	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d -> %s \n", i, slice[i])
	}

}
