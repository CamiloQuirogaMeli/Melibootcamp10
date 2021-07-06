package main

import (
	"fmt"
	"strings"
)

func calcularSueldo(minutos int, categoria string) int {
	categoria = strings.ToUpper(categoria)
	valorHora := map[string]int{"A": 3000, "B": 1500, "C": 1000}
	horas := minutos / 60
	sueldoPorHora := valorHora[categoria] * horas

	if categoria == "A" {
		sueldoPorHora = sueldoPorHora + (50 * sueldoPorHora / 100)
	} else if categoria == "B" {
		sueldoPorHora = sueldoPorHora + (20 * sueldoPorHora / 100)
	}
	return sueldoPorHora
}

func main() {
	sueldo := calcularSueldo(125, "A")
	fmt.Println("Sueldo", sueldo)
}
