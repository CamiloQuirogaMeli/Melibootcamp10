package ejercicios

import "fmt"

func PrimerEjercicio() {
	var salario = 200000
	fmt.Println("El impuesto del salario es ", calcularImpuesto(salario))
}

func calcularImpuesto(salario int) int {
	var porcentaje = 0
	if salario > 50000 {
		porcentaje = 17
		if salario > 150000 {
			porcentaje += 10
		}
	}

	return (salario / 100) * porcentaje
}
