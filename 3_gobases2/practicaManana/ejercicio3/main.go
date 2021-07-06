package main

import (
	"fmt"
)

const CATA, CATB, CATC = "A", "B", "C"

func main() {
	var minutos int
	var categoria string
	var horas int

	fmt.Println("Ingrese categoria del empleado: ")
	fmt.Scanln(&categoria)
	fmt.Println("Ingrese la cantidad de minutos que trabajo: ")
	fmt.Scanln(&minutos)

	horas = redondeoHoras(minutos)

	fmt.Println(calcularSalario(categoria, horas))

}

func redondeoHoras(minutos int) int {
	var horas float32
	horas = float32(minutos) / 60
	return int(horas)

}

func calcularSalario(categoria string, horas int) (salario float32) {
	switch categoria {
	case CATA:
		salario = float32(horas) * 3000
		salario = salario + (salario / 2)

	case CATB:
		salario = float32(horas) * 1500
		salario = salario + (salario * 0.20)

	case CATC:
		salario = float32(horas) * 1000

	}
	return salario
}
