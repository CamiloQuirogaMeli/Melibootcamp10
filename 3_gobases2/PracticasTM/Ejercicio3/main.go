package main

import (
	"fmt"
)

func main() {
	var categoria string = "B" //cambie la categoria
	var minutos float64 = 2000 // cambie los minutos
	imp := salario(minutos, categoria)
	if imp == 0 {
		fmt.Println("No se pudo calcular el salario, no tiene una categoria calculable")
	} else {
		fmt.Println("El salario del trabajor es: ", imp)
	}

}

func salario(minutos float64, categoria string) float64 {
	var salario float64 = 0
	var horas float64 = 0
	var promedio float64
	if categoria == "C" {
		salario = 1000
		horas = minutos / 60
		promedio = salario * horas
	} else if categoria == "B" {
		salario = 1500
		horas = minutos / 60
		promedio = salario * horas
		promedio = promedio + (promedio * 20 / 100)
	} else if categoria == "A" {
		salario = 3000
		horas = minutos / 60
		promedio = salario * horas
		promedio = promedio + (promedio * 50 / 100)
	} else {
		promedio = 0

	}

	return promedio

}
