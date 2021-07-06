package main

import (
	"errors"
	"fmt"
)

type customError struct {
	message string
}

const (
	ValueHour = 500.0
)

func main() {
	salary, err := monthlySalary(90)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Salario total es:", salary)
	}
	fmt.Println("------------------------------------------------")
	bonus, err := bonus(-160000, 8)
	if err != nil {
		fmt.Println(err)
		fmt.Println("el error causa (envuelto) es:", errors.Unwrap(err))
	} else {
		fmt.Printf("El aguinaldo sera de: %.2f\n\n", bonus)
	}
}

func (err *customError) Error() string{
	return err.message
}

func monthlySalary(hours int) (float64, error){

	if hours < 80 {
		err := errors.New("error: el trabajador no puede haber trabajado menos de 80hs mensuales")
		return -1.0, err
	}
	salary := float64(hours) * ValueHour
	if  salary > 150000.0 {
		salary *= 0.9
	}
	return salary, nil

}

func bonus(bestSalary float64, monthsActive int) (float64, error){
	var bonus float64
	bonus = (bestSalary / 12) * float64(monthsActive)
	if bestSalary < 0 || monthsActive < 0 {
		err := &customError{message: "ingreso un valor negativo"}
		err2 := fmt.Errorf("error: %w", err)
		return -1.0, err2
	}
	return bonus, nil
}