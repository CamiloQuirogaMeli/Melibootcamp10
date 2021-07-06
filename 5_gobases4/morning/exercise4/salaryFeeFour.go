package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	MIN_SALARY     = 150000
	FEE_PERCENTAGE = 0.1
)

func main() {

	workedHours := []int{160, 200, 87, 90, 170, -1}
	hourValue := 45000
	agi, err := getAguinaldo(workedHours, hourValue)
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))
	} else {
		fmt.Println("Su aguinaldo semestral es de", agi)
	}

}

func getNetSalary(hours int, hourValue int) (int, error) {
	if hours < 80 {
		return 0, errors.New("el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	bruteSalary := float64(hourValue * hours)
	if bruteSalary >= MIN_SALARY {
		return int(bruteSalary * (1 - FEE_PERCENTAGE)), nil
	} else {
		return int(bruteSalary), nil
	}
}

func getAguinaldo(hoursWorked []int, hourValue int) (int, error) {
	if len(hoursWorked) > 6 {
		return 0, errors.New("el trabajador no puede trabajar más de 6 meses en un semestre")
	}
	if len(hoursWorked) == 0 {
		return 0, nil
	}

	bestSalary := math.Inf(-1)
	for _, monthHours := range hoursWorked {
		netSalary, err := getNetSalary(monthHours, hourValue)
		if err != nil {
			return 0, fmt.Errorf("error en aguinaldo: %w", err)
		} else {
			fmt.Println("Su salario este mes es de", netSalary)
			bestSalary = math.Max(float64(netSalary), bestSalary)
		}
	}
	return (int(bestSalary/12) * len(hoursWorked)), nil
}
