package main

import (
	"errors"
	"fmt"
)

func calculoHoras(minutos int) int {
	return minutos / 60
}

func calculoSalario(categoria string, minutos int) (float64, error) {
	switch categoria {
	case "C":
		{
			return float64(calculoHoras(minutos) * 1000), nil
		}
	case "B":
		{
			horas := calculoHoras(minutos)
			return float64(horas)*1500 + (float64(horas) * 1500 * 0.20), nil
		}
	case "A":
		{
			horas := calculoHoras(minutos)
			return float64(horas)*3000 + (float64(horas) * 3000 * 0.50), nil
		}
	default:
		{
			error := errors.New("la categoria no existe")
			return 0, error
		}
	}
}

func main() {
	var (
		categoria string
		minutos   int
	)
	fmt.Println("Ingrese su categoria")
	fmt.Scanln(&categoria)
	fmt.Println("Ingrese los minutos trabajados en el mes")
	fmt.Scanln(&minutos)
	res, err := calculoSalario(categoria, minutos)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Su salario del mes es ", res)
	}
}
