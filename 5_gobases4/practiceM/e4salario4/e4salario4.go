package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	pricePerHours := -42.3
	hours := 89
	salary, err := calculateSalary(hours, pricePerHours)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bonus, err := calculateBonus(salary) // Acá supongo que este salary es el mejor salario del semestre.
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("EL aguinaldo es: %.2f\n", bonus)
}

func calculateBonus(salary float64) (float64, error) {
	if salary < 0 {
		return 0, errors.New(fmt.Sprintf("%.2f - error: salario no puede ser negativo", salary))
	} else {
		return (salary / 12) * 6, nil
	}
}

func calculateSalary(hours int, pricePerHours float64) (float64, error) {
	switch {
	case hours < 0:
		return 0, errors.New(fmt.Sprintf("%d - %.2f error: cantidad de horas inválida.", hours, pricePerHours))
	case hours < 80:
		return 0, errors.New(fmt.Sprintf("%d - %.2f error: el trabajadorno pudo haber trabajado menos de 80 horas mensuales.", hours, pricePerHours))
	case pricePerHours < 0:
		return 0, errors.New(fmt.Sprintf("%d - %.2f error: salario inválido", hours, pricePerHours))
	default:
		return (pricePerHours * float64(hours)), nil
	}
}
