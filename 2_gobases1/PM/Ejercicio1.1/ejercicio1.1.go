package main

import (
	"fmt"
	"strings"
)

func main() {
	var palabra string = "Aquelarre"
	newArr := strings.Split(palabra, "")

	fmt.Println(len(newArr))

	for i := 0; i < len(newArr); i++ {
		fmt.Printf("%d %v\n", i+1, newArr[i])
	}

}
