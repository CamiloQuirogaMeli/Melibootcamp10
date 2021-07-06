package main

import (
	"errors"
	"fmt"
)

func sueldoEmpresaMarinera(cat string, min uint) (float64, error) {
	switch cat {
	case "A":
		sueldoHora := 3000.0
		horas := min / 60
		salarioMensual := sueldoHora * float64(horas)
		return salarioMensual * 1.5, nil

	case "B":
		sueldoHora := 1500.0
		horas := min / 60
		salarioMensual := sueldoHora * float64(horas)
		return salarioMensual * 1.2, nil

	case "C":
		sueldoHora := 1000.0
		horas := min / 60
		salarioMensual := sueldoHora * float64(horas)
		return salarioMensual, nil
	default:
		return 0, errors.New("categoria incorrecta")
	}

}

func main() {
	res, err := sueldoEmpresaMarinera("A", 120)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario del trabajador es: %.2f\n", res)
	}

}
