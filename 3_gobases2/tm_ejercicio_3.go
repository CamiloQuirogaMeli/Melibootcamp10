package main

import (
	"errors"
	"fmt"
)

func main() {
	categoria, minutos := "D", 600.0

	salario, err := calcularSalario(minutos, categoria)

	if salario < 0 {
		fmt.Printf("%s \n", err)
	} else {
		fmt.Printf("El salario calculado es: %f \n", salario)
	}

}

func calcularSalario(minutos float64, categoria string) (float64, error) {

	horas := minutos / 60

	switch categoria {
	case "A":
		return horas * 3000 * 1.5, errors.New("")
	case "B":
		return horas * 1500 * 1.2, errors.New("")
	case "C":
		return horas * 1000, errors.New("")
	default:
		return -1, errors.New("Se ingreso una categoria inexistente")
	}
}
