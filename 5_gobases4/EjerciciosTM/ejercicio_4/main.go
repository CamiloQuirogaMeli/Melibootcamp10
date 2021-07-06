package main

import (
	"errors"
	"fmt"
)

const (
	minimoImponible = 150000
	precioPorHora   = 200
)

func calcularSalario(horasTrabajadas int) (float64, error) {
	if horasTrabajadas < 80 {
		return 0, fmt.Errorf("el trabajador no puede haber trabajado menos de 80 hs mensuales. Horas ingresadas %d:", horasTrabajadas)
	}

	salario := float64(horasTrabajadas * precioPorHora)

	if salario < minimoImponible {
		return salario, nil
	}

	return salario - (salario * 0.1), nil
}

func calcularMedioAguinaldo(mejorSalario float64, mesesTrabajados int) (float64, error) {
	if mejorSalario < 0 {
		return 0, errors.New("error: se ingreso el mejor salario negativo")
	}

	return (mejorSalario / 12) * float64(mesesTrabajados), nil
}

func main() {

	fmt.Println("Ingrese horas trabajadas:")
	var horasTrabajadas int
	fmt.Scanln(&horasTrabajadas)

	fmt.Println("Ingrese su mejor salario:")
	var mejorSalario float64
	fmt.Scanln(&mejorSalario)

	fmt.Println("Ingrese los meses trabajados:")
	var mesesTrabajados int
	fmt.Scanln(&mesesTrabajados)

	salario, err := calcularSalario(horasTrabajadas)

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("salario: ", salario)

	aguinaldo, err := calcularMedioAguinaldo(mejorSalario, mesesTrabajados)

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("medio aguinaldo: ", aguinaldo)

}
