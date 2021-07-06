package main

import (
	"errors"
	"fmt"
)

const CAT_A = "A"
const CAT_B = "B"
const CAT_C = "C"

func calcularSalarioMensual(minutosPorMes int, categoria string) (result int, err error) {

	horasPorMes := minutosPorMes / 60

	switch categoria {
	case CAT_C:
		result = 1000 * horasPorMes
		err = nil

	case CAT_B:
		result = 1500 * horasPorMes
		salarioMensualPorcent := float32(result) * 0.2
		result = result + int(salarioMensualPorcent)
		err = nil

	case CAT_A:
		result = 3000 * horasPorMes
		salarioMensualPorcent := float32(result) * 0.2
		result = result + int(salarioMensualPorcent)
		err = nil

	default:
		err = errors.New("categoria incorrecta")
		result = 0
	}

	return
}

func main() {

	salario, err := calcularSalarioMensual(200, CAT_A)
	if err == nil {
		fmt.Println("Salario: ", salario)
	} else {
		fmt.Println("Error: ", err)
	}

}
