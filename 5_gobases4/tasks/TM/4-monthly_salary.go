package main

import (
	"errors"
)

func monthlySalary(hours int) (netSalary float64, err error) {
	baseHourlySalary := 2000

	if hours < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	netSalary = float64(hours) * float64(baseHourlySalary)

	if netSalary > 150000 {
		deduction := netSalary * 0.10
		netSalary -= deduction
	}

	return netSalary, nil
}
