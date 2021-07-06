package main

import "fmt"

func main() {
	minFunc := calcularEstadisticas("minimo")
	maxFunc := calcularEstadisticas("maximo")
	promFunc := calcularEstadisticas("promedio")
	errFunc := calcularEstadisticas("prdio")

	valMinimo := minFunc(4, 2.3, 5, 2.1, 4, 4.3, 4.8)
	valMaximo := maxFunc(4, 2.3, 5, 2.1, 4, 4.3, 4.8)
	valPromedio := promFunc(4, 2.3, 5, 2.1, 4, 4.3, 4.8)

	fmt.Println(valMinimo, valMaximo, valPromedio)
}
