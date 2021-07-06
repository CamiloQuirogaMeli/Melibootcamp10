package main

import "fmt"

func devolverImpuesto(s float32) float32 {
	if s > 150000 {
		return (10 * s) / 100
	} else if s > 50000 {
		return (17 * s) / 100
	} else {
		return 0
	}
}

func main() {
	var sueldo float32

	fmt.Printf("Ingrese el salario: ")
	fmt.Scanf("%f", &sueldo)

	imp := devolverImpuesto(sueldo)

	fmt.Printf("Salario: %.2f", imp+sueldo)
}
