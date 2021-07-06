package main

import (
	"errors"
	"fmt"
	"os"
)

func salarioMensual(horasTrabajadas int) (int, error) {
	precioHora := 70
	if horasTrabajadas < 80 {
		return 0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales. Horas trabajadas %d", horasTrabajadas)
	} else {
		salario := precioHora * horasTrabajadas
		descuento := salario * 10 / 100
		return salario - descuento, nil
	}
}

func medioAguinaldo(mejorSalario int, mesesTrabajados int) (int, error) {
	if mejorSalario <= 0 || mesesTrabajados <= 0 {
		return 0, errors.New("error: Ingrese un valor valido")
	} else {
		return mejorSalario / 12 * mesesTrabajados, nil
	}
}

func main() {
	salario, err := salarioMensual(90)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Salario Mensual:", salario)

	aguinaldo, err := medioAguinaldo(salario, 6)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Medio Aguinaldo:", aguinaldo)
}
