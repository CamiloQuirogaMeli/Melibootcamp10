package main

import (
	"errors"
	"fmt"
)

const salaryA float32 = 1000
const salaryB float32 = 1500
const salaryC float32 = 3000
const extraB float32 = 1.2
const extraC float32 = 1.5

func main() {

	category, min := dataEntry()
	salary, err := salCalc(category, min)

	if err != nil {
		fmt.Print(err)
	} else {

		fmt.Printf("El salario es %.2f\n", salary)
	}
}

func dataEntry() (int, int) {
	var min int
	var category int
	fmt.Println("Ingrese la cantidad de minutos trabajados:")
	fmt.Scanln(&min)

	fmt.Printf("Ingrese la categoria del empleado: \n")
	fmt.Printf("1 - Categoría A \n")
	fmt.Printf("2 - Categoría B \n")
	fmt.Printf("3 - Categoría C \n")
	fmt.Scanln(&category)

	return category, min
}

func salCalc(category, min int) (float32, error) {
	hours := float32(min) / 60
	switch category {
	case 1:
		return float32(hours) * salaryA, nil
	case 2:
		return float32(hours) * salaryB * extraB, nil
	case 3:
		return float32(hours) * salaryC * extraC, nil
	}
	return 0.0, errors.New("La categoria es incorrecta\n")
}
