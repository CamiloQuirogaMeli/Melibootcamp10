package main

import (
	"fmt"
	"errors"
)

const (
	MINIMO = "minimo"
	PROMEDIO = "promedio"
	MAXIMO = "maximo"
)

func max(values... int) int {
	maxValue := values[0]
	for _, value := range values {
		if maxValue < value {
			maxValue = value
		}
	}

	return maxValue
}

func min(values... int) int {
	minValue := values[0]
	for _, value := range values {
		if minValue > value {
			minValue = value
		}
	}

	return minValue
}

func avg(values... int) int {
	var total int = 0
	for _, value := range values {
		total += value
	}
	return total / len(values)
}

func operacion(operationType string) (func(...int) int, error) {
	switch operationType {
		case MINIMO: 
			return min, nil
		case PROMEDIO: 
			return avg, nil
		case MAXIMO: 
			return max, nil
	}
	return nil, errors.New("La operacion no esta definida")
}

func main() {
	stadistics := []int{2,3,3,4,1,2,4,5}

	minFunc, _ := operacion(MINIMO)
	valorMinimo := minFunc(2,3,3,4,1,2,4,5)
	fmt.Println("* Caso funcion minimo")
	fmt.Println("\tEstadisticas:", stadistics)
	fmt.Println("\tMinimo:", valorMinimo)

	maxFunc, _ := operacion(MAXIMO)
	valorMaximo := maxFunc(2,3,3,4,1,2,4,5)
	fmt.Println("* Caso funcion maximo")
	fmt.Println("\tEstadisticas:", stadistics)
	fmt.Println("\tMaximo:", valorMaximo)

	promFunc, _ := operacion(PROMEDIO)
	valorPromedio := promFunc(2,3,3,4,1,2,4,5)
	fmt.Println("* Caso funcion promedio")
	fmt.Println("\tEstadisticas:", stadistics)
	fmt.Println("\tPromedio:", valorPromedio)

	_, err := operacion("abc")
	fmt.Println("* Caso funcion inexistente")
	fmt.Println("\tError:", err)
}
