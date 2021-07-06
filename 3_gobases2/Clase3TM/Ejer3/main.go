package main

import (
	"errors"
	"fmt"
)

func main() {
	salary, err := salary(60, "C", -100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario es de ", salary)
	}
}
func salary(minutes int, category string, monthlySalary float64) (float64, error) {
	if minutes < 0 {
		return 0, errors.New("No se puede haber trabajado una cantidad negativa de minutos")
	}
	if monthlySalary < 0 {
		return 0, errors.New("El salario mensual no puede ser negativo")
	}
	switch category {
	case "A":
		return float64(minutes) / 60.0 * 1000.0, nil
	case "B":
		return float64(minutes)/60.0*1500.0 + monthlySalary*0.2, nil
	case "C":
		return float64(minutes)/60.0*3000.0 + monthlySalary*0.5, nil
	default:
		return 0, errors.New("Categoria invalida")
	}
}
