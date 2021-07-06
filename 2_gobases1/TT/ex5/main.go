package main

import (
	"fmt"
)

func main() {
	var list = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println("Lista:")
	for i := 0; i < len(list); i++ {
		fmt.Println(" *", list[i])
	}
	fmt.Println()
	list = append(list, "Gabriela")
	fmt.Println("Nueva lista:")
	for i := 0; i < len(list); i++ {
		fmt.Println(" *", list[i])
	}
}
