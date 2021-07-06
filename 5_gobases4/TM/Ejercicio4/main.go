package main

import (
	"errors"
	"fmt"
)

type errorPersonalizado struct{}

func (e *errorPersonalizado) Error() string {
	return "error: el número no puede ser negativo"
}

func calcularSalario(horas int, valorHora float64) (salario float64, e error) {
	if horas < 80 {
		e = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	} else if valorHora < 1 {
		e = fmt.Errorf("error: el valor hora no puede ser menor a 1 e ingresó: %v", valorHora)
	} else {
		if valorHora*float64(horas) >= 150000 {
			salario = (valorHora * float64(horas)) * 0.9
		} else {
			salario = valorHora * float64(horas)
		}
	}
	return
}

func calcularAguinaldo(meses int, mejorSalario float64) (aguinaldo float64, e error) {
	e1 := errorPersonalizado{}
	if mejorSalario < 0 || meses < 0 {
		e = fmt.Errorf("%w", &e1)
	} else {
		aguinaldo = (mejorSalario / 12) * float64(meses)
	}
	return
}

func main() {

	salario, err := calcularSalario(70, 800)
	if err == nil {
		fmt.Println("El salario es de: ", salario)
	} else {
		fmt.Println(err)
	}

	salario1, err1 := calcularSalario(150, 765)
	if err1 == nil {
		fmt.Println("El salario es de: ", salario1)
	} else {
		fmt.Println(err1)
	}

	aguinaldo, err3 := calcularAguinaldo(11, 154000)
	if err3 == nil {
		fmt.Printf("El aguinaldo es de: %.2f \n", aguinaldo)
	} else {
		fmt.Println(err3)
	}

	aguinaldo1, err4 := calcularAguinaldo(-11, 155555)
	if err4 == nil {
		fmt.Printf("El aguinaldo es de: %.2f \n", aguinaldo1)
	} else {
		fmt.Println(err4)
	}
}
