package ejercicios

import (
	"errors"
	"fmt"
)

func TercerEjercicio() {
	salario, err := calcularSalario(1800, "A")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("El salario del empleado es: %5.2f\n", salario)

	salario, err = calcularSalario(1800, "B")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("El salario del empleado es: %5.2f\n", salario)

	salario, err = calcularSalario(1800, "C")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("El salario del empleado es: %.2f\n", salario)
}

func calcularSalario(minutosTrabajados int, categoria string) (float32, error) {
	var horasTrabajadas = minutosTrabajados / 60
	var salarioHora = 0
	var porcentajeAdicional float32

	switch categoria {
	case "A":
		salarioHora = 1000
		porcentajeAdicional = 0

	case "B":
		salarioHora = 1500
		porcentajeAdicional = 0.2

	case "C":
		salarioHora = 3000
		porcentajeAdicional = 0.5

	default:
		return 0, errors.New("categoria no válida")
	}

	var totalSalario = float32(salarioHora) * float32(horasTrabajadas)
	return totalSalario + totalSalario*porcentajeAdicional, nil
}
