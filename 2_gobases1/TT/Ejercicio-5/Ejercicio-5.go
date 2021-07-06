package main

import "fmt"

func main() {
	fmt.Println("Ejercicio 5")
	nameList := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}
	nameList = append(nameList, "Gabriela")
	for _, name := range nameList {
		fmt.Println(name)
	}
}
