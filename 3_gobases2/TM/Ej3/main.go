package main

import "fmt"

const (
	A = "a"
	B = "b"
	C = "c"
)

func main() {
	fmt.Printf("Empleado 1 \nCategoria: C\nHoras trabajadas: 50\nSalario: %.2f\n\n", calcularSalario(50, C))
	fmt.Printf("Empleado 2 \nCategoria: B\nHoras trabajadas: 70\nSalario: %.2f\n\n", calcularSalario(70, B))
	fmt.Printf("Empleado 3 \nCategoria: A\nHoras trabajadas: 60\nSalario: %.2f\n\n", calcularSalario(60, A))
}

func calcularSalario(horas int, categoria string) float64 {
	switch categoria {
	case C:
		return float64(horas) * 1000
	case B:
		return (float64(horas) * 1500) * 1.2
	case A:
		return (float64(horas) * 3000) * 1.5
	default:
		return 0
	}
}
