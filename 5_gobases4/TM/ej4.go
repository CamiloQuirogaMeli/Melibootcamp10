package main

import (
	"errors"
	"fmt"
)

func salarioMensual(salarioMensual float64, horasTrabajadas int) (float64, error) {
	if horasTrabajadas < 80 {
		e := errors.New("error: el trabajador no puede haber trabajado menos de 80hs mensuales")
		return 0, e
	}

	if salarioMensual >= 150000 {
		e := errors.New("error: el salario ingresado no alcanza el minimo imponible")
		s := float64(salarioMensual) * 0.9
		return s, e
	}

	return float64(salarioMensual), nil
}

func calcularAguinaldo(salariosSemestre [6]float64, mesesTrabajados int) (float64, error) {
	if mesesTrabajados < 0 {
		return 0, fmt.Errorf("error: el valor de meses trabajados no puede ser negativo: %d", mesesTrabajados)
	}
	max := 0.0
	for _, valor := range salariosSemestre {
		if valor > max {
			max = valor
		}
	}
	salario := (max / 12.0) * float64(mesesTrabajados)
	return float64(salario), nil
}

func main() {

	salario := 160000.0

	sm, err := salarioMensual(salario, 90)

	fmt.Println("SM", sm)
	fmt.Println("err", err)

	salariosSemestre := [6]float64{140000.0, 120000.0, 120000.0, 120000.0, 130000.0, 180000.0}
	sm1, err1 := calcularAguinaldo(salariosSemestre, 6)

	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Println("su medio aguinaldo es: ", sm1)

	e := fmt.Errorf("%w and %w", err, err1)

	var errr error
	errr = errors.Unwrap(e)
	fmt.Println("error 1 unwrap", errr)
	err3 := errors.Unwrap(errr)
	fmt.Println("error 2 unwrap", err3)
}
