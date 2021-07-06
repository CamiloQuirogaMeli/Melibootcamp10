package ejercicios

import (
	"errors"
	"fmt"
)

func SegundoEjercicio() {
	var notas = []int{5, 6, 1, 8, 9, 3, 6, 7}
	promedio, err := calcularPromedioNotas(notas)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio de notas del curso es %d\n", promedio)
	}
}

func calcularPromedioNotas(notas []int) (int, error) {
	var sum = 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("nota no puede ser negativa")
		} else {
			sum += nota
		}
	}

	return sum / len(notas), nil
}
