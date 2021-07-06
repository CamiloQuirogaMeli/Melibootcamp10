package main

import (
	"errors"
	"fmt"
)

func calculoSalario(salary float64, horas int) (float64, error) {
	switch {
	case salary > 0 && horas > 80:
		{
			salario := salary * float64(horas)
			if salario > 150000 {
				salario = salario + (salario * 0.10)
			}
			return salario, nil
		}
	case horas < 80:
		{
			return 0, errors.New("el trabajador no puede haber trabajado menos de 80hs mensuales")
		}
	default:
		{
			return 0, fmt.Errorf("el trabajador no puede tener un valor hora menor a 1, y su valor hora es %f", salary)
		}
	}
}

func calculoAguinaldo(mesesTrabajados int, salarios ...float64) (float64, error) {
	var aguinaldo float64
	mejorSalario := salarios[0]
	for _, salario := range salarios {
		if salario < 0 {
			return 0, fmt.Errorf("no puede tener un salario negativo")
		}
		if mejorSalario < salario {
			mejorSalario = salario
		}
	}
	if mesesTrabajados < 1 {
		return 0, errors.New("error tiene que haber trabajado al menos un mes en el semestre")
	}
	aguinaldo = mejorSalario / 12 * float64(mesesTrabajados)
	return aguinaldo, nil
}

func main() {
	var (
		salary float64
		horas  int
	)
	fmt.Println("Ingrese su valor x hora")
	fmt.Scanln(&salary)
	fmt.Println("Ingrese las horas trabajadas en el mes")
	fmt.Scanln(&horas)
	mensual, err := calculoSalario(salary, horas)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Printf("El salario caldulado es de %.2f", mensual)
		fmt.Println("Se calculara el aguinaldo")
		aguinaldo, err := calculoAguinaldo(5, mensual, 20000, 3000000, 400000)
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Printf("el aguinaldo calculado es de %.2f", aguinaldo)
		}
	}
}
