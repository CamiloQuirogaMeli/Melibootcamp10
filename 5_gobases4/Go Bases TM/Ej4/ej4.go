package main

import (
	"errors"
	"fmt"
	"os"
)

type NumeroNegativoError struct{}

func (n *NumeroNegativoError) Error() string {
	return "El valor ingresado es negativo"
}

func main() {
	cantHoras, opcion := 0, -1
	for opcion != 0 {
		fmt.Println("0. Salir")
		fmt.Println("1. Calcular sueldo mensual")
		fmt.Println("2. Calcu")
		fmt.Scanln(&opcion)
		switch opcion {
		case 0:
			os.Exit(1)
		case 1:
			fmt.Println("Cantidad de horas: ")
			fmt.Scanln(&cantHoras)
			salario, err := salarioMensual(cantHoras)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Su salario mensual es de: ", salario)
		case 2:
			valorAguinaldo, err := medioAguinaldo()
			if err != nil {
				fmt.Println(errors.Unwrap(err))
				continue
			}
			fmt.Println("El medio aguinaldo será de ", valorAguinaldo)
		}
	}
}

func salarioMensual(cantHoras int) (float32, error) {
	if cantHoras < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	salario := float32(cantHoras * 1300)
	if salario >= 150000 {
		return salario - (salario * 0.1), nil
	} else {
		return salario, nil
	}
}
func medioAguinaldo() (float64, error) {
	mejorSalario, cantMeses := 0.0, 0
	fmt.Println("Mejor salario del semestre: ")
	fmt.Scanln(&mejorSalario)
	if mejorSalario < 0 {
		return 0.0, fmt.Errorf("Error: %w", &NumeroNegativoError{})
	}
	fmt.Println("Cantidad de meses trabajados en el semestre: ")
	fmt.Scanln(&cantMeses)
	if cantMeses < 0 {
		e := NumeroNegativoError{}
		return 0.0, fmt.Errorf("Error: %w", &e)
	}
	return (mejorSalario / 12) * float64(cantMeses), nil
}
