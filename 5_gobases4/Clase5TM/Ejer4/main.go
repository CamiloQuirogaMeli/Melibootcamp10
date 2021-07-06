package main

import (
	"errors"
	"fmt"
)

func main() {
	salarySemester := []float64{100000, 200000, 300000, 400000, 500000, 600000}

	ms, err := monthlySalary(80, salarySemester[2])

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario final es de %.2f\n", ms)
	}

	bb, err := bestBonus(salarySemester, 6)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El aguinaldo es de %.2f\n", bb)
	}
}

func monthlySalary(hours int, salary float64) (float64, error) {
	if hours < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	} else {
		if salary < 0 {
			return 0, errors.New("error: el salario no puede ser negativo")
		} else if salary >= 150000 {
			return salary * 0.9, nil
		} else {
			return salary, nil
		}
	}
}

func bestBonus(salarySemester []float64, monthsWorked int) (float64, error) {
	var max float64

	for _, motnhValue := range salarySemester {
		if motnhValue < 0 {
			return 0, fmt.Errorf("error: no puede tener meses con sueldo negativo")
		}
		if max < motnhValue {
			max = motnhValue
		}
	}

	if monthsWorked < 1 {
		return 0, fmt.Errorf("error: los meses trabajados no pueden ser menores o iguales a 0")
	}

	return max / 12 * float64(monthsWorked), nil
}
