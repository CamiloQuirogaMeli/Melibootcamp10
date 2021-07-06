package main

import (
	"fmt"
	"strings"
)

func main() {
	var hours int
	fmt.Print("Escriba la cantidad de horas trabajadas: ")
	fmt.Scanln(&hours)
	var cat string
	fmt.Print("Escriba la categoría del trabajador (A, B o C): ")
	fmt.Scanln(&cat)
	cat = strings.ToLower(cat)
	var salary float64
	switch cat {
	case "a":
		salary = float64(hours) * 1000.0
	case "b":
		salary = float64(hours) * 1500.0
		salary *= 1.2
	case "c":
		salary = float64(hours) * 3000.0
		salary *= 1.5
	}
	fmt.Println("El salario es: ", salary)
}
