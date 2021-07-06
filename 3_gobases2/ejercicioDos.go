package main

import "fmt"

func calcularPromedioNotas(notas ...int) int {
	var promedio int = 0
	for _, value := range notas {
		promedio += value
	}
	return promedio / len(notas)
}

func main() {
	fmt.Printf("EL promedio del alumno es de: %d\n", calcularPromedioNotas(4, 5, 3, 2))
}
