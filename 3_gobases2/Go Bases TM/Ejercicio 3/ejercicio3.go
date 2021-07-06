package main

import (
	"errors"
	"fmt"
)

func main() {
	var minutos int
	fmt.Println("Cantidad de minutos trabajados: ")
	fmt.Scanln(&minutos)

	var categoria string
	fmt.Println("Categoría: ")
	fmt.Scanln(&categoria)

	salario, err := calcularSueldo(minutos, categoria)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Su sueldo es de: ", salario)
	}

}

func calcularSueldo(minutos int, categoria string) (salario float32, err error) {
	if minutos <= 0 {
		return 0.0, errors.New("Cantidad de minutos inválida")
	}
	var sueldo float32
	switch categoria {
	case "A":
		sueldo = float32(minutos * 3000)
		return sueldo + sueldo*.5, nil
	case "B":
		sueldo = float32(minutos * 1500)
		return sueldo + sueldo*.2, nil
	case "C":
		return float32(minutos * 1000), nil
	default:
		return 0, errors.New("Categoría inválida")
	}
}
