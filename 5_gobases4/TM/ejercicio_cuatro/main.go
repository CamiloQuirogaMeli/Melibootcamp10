package main

import (
	"errors"
	"fmt"
)

const (
	COSTO_HORA          = 1500
	TOPE                = 150000
	CANT_MESES_SEMESTRE = -6
)

func calcularMedioAginaldo(salarioMensual float32, cantidadMesesSemestre int) (medioAginaldo float32, err error) {

	err = nil
	medioAginaldo = 0

	if salarioMensual < 0 || cantidadMesesSemestre < 0 {
		err = fmt.Errorf("error: no puede ingresar un numero negativo. salario: %.2f cantidad horas semestre: %d", salarioMensual, cantidadMesesSemestre)
		return
	}

	medioAginaldo = (salarioMensual / 12) * float32(cantidadMesesSemestre)
	return
}

func salario(cantHoras int) (salarioMensual float32, medioAginaldo float32, err error) {

	err = nil
	salarioMensual = 0
	medioAginaldo = 0
	if cantHoras < 80 {
		err = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return
	}

	salarioMensual = float32(COSTO_HORA * cantHoras)
	if salarioMensual > TOPE {
		salarioMensual = salarioMensual * 0.9
	}

	medioAginaldo, medioAginaldoErr := calcularMedioAginaldo(salarioMensual, CANT_MESES_SEMESTRE)

	if medioAginaldoErr != nil {
		err = fmt.Errorf("Se detecto un error al calcular el medio aginaldo %w", medioAginaldoErr)
	}

	return
}

func main() {

	salarioMensual, medioAginaldo, err := salario(90)

	if err != nil {

		error_dos := errors.Unwrap(err)
		if error_dos != nil {
			fmt.Println(error_dos)
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Printf("El salario mensual es de %.2f con un medio aginaldo de %.2f", salarioMensual, medioAginaldo)
	}

}
