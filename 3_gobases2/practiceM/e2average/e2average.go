package main

import "fmt"

func main() {
	// Ingrese nombre del alumno ingrese calificacion o OK para continuar
	aux := average(10, 10, 5, 5)
	fmt.Println("El promedio es: ", aux)
}

func average(qualifi ...float64) float64 {
	var count float64
	for _, val := range qualifi {
		count = count + val
	}
	return count / float64(len(qualifi))
}
