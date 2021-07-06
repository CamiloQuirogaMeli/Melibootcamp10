package main

import (
	"errors"
	"fmt"
)

type errorValorNegativo struct{}

func (e errorValorNegativo) Error() string {
	return "valor negativo ingresado"
}

func calcularMedioAguinaldo(mejorSalario float64, meses int) (float64, error) {
	if mejorSalario > 0 && meses > 0 {
		return (mejorSalario / 12) * float64(meses), nil
	}
	return 0, errorValorNegativo{}
}

func calcularSalarioMensual(valorHora float64, horasMes int) (float64, error) {
	if horasMes > 80 {
		bruto := valorHora * float64(horasMes)
		if bruto > 150000 {
			return bruto * 0.9, nil
		}
		return bruto, nil
	}
	return 0, errors.New("el trabajador no puede haber trabajado menos de 80 horas mensuales")

}

func main() {
	//Item 1
	/*salarioMensual := 160000
	horasMes := 4

	result, err := calcularSalarioMensual(float64(salarioMensual), horasMes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("El salario mensual neto es de ", result)*/

	//Item 2
	mejorSalario := -30
	meses := 4

	result, err := calcularMedioAguinaldo(float64(mejorSalario), meses)
	if err != nil {
		err2 := fmt.Errorf("error: %w", err)
		fmt.Println(errors.Unwrap(err2))
		return
	}

	fmt.Println("El medio aguinaldo es de: ", result)

}
