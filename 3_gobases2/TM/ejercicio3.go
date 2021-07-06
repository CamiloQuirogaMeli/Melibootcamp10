package main

import (
	"errors"
	"fmt"
)

func main() {
	salary, err := getSalary(2, "X")
	if err == nil {
		fmt.Printf("El salario es: $%.2f\n", salary)
	} else {
		fmt.Println(err)
	}
}

func getSalary(hours float64, category string) (float64, error) {
	switch category {
	case "C":
		return 1000 * hours, nil
	case "B":
		return (1500 * hours) * 1.2, nil
	case "A":
		return (3000 * hours) * 1.5, nil
	default:
		return 0, errors.New("Categoria invalida")
	}
}
