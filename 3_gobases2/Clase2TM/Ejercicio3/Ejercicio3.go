package main

import (
	"fmt"
)

func main() {

	var cat string
	var minutes int
	fmt.Println("Ingrese la categoria A, B o C: ")
	fmt.Scanf("%s", &cat)
	fmt.Println("Ingrese la cantidad de minutos trabajados: ")
	fmt.Scanf("%d", &minutes)
	if (cat == "A" || cat == "a" || cat == "B" || cat == "b" || cat == "C" || cat == "c") && (minutes > 0) {
		sal := salary(minutes, cat)
		fmt.Printf("El sueldo es de: %d", sal)
	} else {
		fmt.Printf("Los minutos ingresados o la categoria no son correctos")
	}
}

func salary(minutes int, category string) int {
	var sal int
	switch {
	case category == "a" || category == "A":
		sal = (minutes / 60) * 3000
		sal = sal + ((sal * 50) / 100)
	case category == "b" || category == "B":
		sal = (minutes / 60) * 1500
		sal = sal + ((sal * 20) / 100)
	case category == "c" || category == "C":
		sal = (minutes / 60) * 1000
	}

	return sal
}
