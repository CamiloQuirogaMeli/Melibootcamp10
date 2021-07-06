package main

import (
	"errors"
	"fmt"
)

type Trabajador struct {
	salarioPorHora      float64
	horasTrabajadas     int
	salarioUltimosMeses []float64
}

func salarioMensual(t Trabajador) (float64, error) {
	if t.horasTrabajadas < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	return float64(t.horasTrabajadas) * t.salarioPorHora, nil
}

func medioAguinaldo(t Trabajador) (float64, error) {
	aguinaldo, err := encontrarSalarioMaximo(t.salarioUltimosMeses)
	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}
	aguinaldo = aguinaldo / 12 * float64(len(t.salarioUltimosMeses))
	return aguinaldo, nil
}

func encontrarSalarioMaximo(arreglo []float64) (float64, error) {
	var maximo float64
	if len(arreglo) > 0 {
		maximo = arreglo[0]
	}
	for _, valor := range arreglo {
		if valor < 0 {
			return 0, fmt.Errorf("no puede haber un salario negativo")
		}
		if valor > maximo {
			maximo = valor
		}
	}
	return maximo, nil
}

func main() {
	trabajador := Trabajador{
		salarioPorHora:      800,
		horasTrabajadas:     100,
		salarioUltimosMeses: []float64{70000, -1, 75000, 70000, 75000, 75000},
	}
	salario, err1 := salarioMensual(trabajador)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("El salario mensual es:", salario)
	}
	aguinaldo, err2 := medioAguinaldo(trabajador)
	if err2 != nil {
		fmt.Println(errors.Unwrap(err2))
	} else {
		fmt.Println("El medio aguinaldo es:", aguinaldo)
	}
}
