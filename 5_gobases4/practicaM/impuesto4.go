package main

import (
	"errors"
	"fmt"
	"os"
)

func salarioTotal(horas int) (int, error) {
	if horas < 80 || horas < 0 {
		return 0, errors.New("error salarioTotal: el trabajador no puede haber trabajado menos de 80 hs mensuales.")
	}

	salario := horas * 4000
	if salario >= 150000 {
		descuento := salario * (10 / 100)
		salario -= descuento
	}
	return salario, nil
}

func medioAguinaldo(salario []int) (int, error) {
	mejorSalario := salario[0]
	for i := range salario {
		if salario[i] < 0 {
			err := fmt.Errorf("ningun salario puede ser negativo: %d es negativo", salario[i])
			return 0, err
		}

		var actual int = salario[i]
		if actual > mejorSalario {
			mejorSalario = actual
		}
	}

	aguinaldo := mejorSalario / 12
	aguinaldo *= len(salario)

	return aguinaldo, nil
}

func main() {
	salary, err2 := salarioTotal(20)

	err1 := fmt.Errorf("err2: %w", err2)
	var salaries = []int{80000, -100000, 150000, 100000}
	aguinaldo, err2 := medioAguinaldo(salaries)

	if err2 != nil {
		fmt.Printf("%s \n", err2)
		fmt.Println(errors.Unwrap(err1))
		os.Exit(1)
	}

	fmt.Printf("Salario total de este mes: %d Valor aguinaldo: %d", salary, aguinaldo)
}
