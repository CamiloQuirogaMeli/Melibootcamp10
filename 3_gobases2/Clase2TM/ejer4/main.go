package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	MINIMO   = "minimo"
	PROMEDIO = "promedio"
	MAXIMO   = "maximo"
)

func getMinimum(scores ...int) int {

	var minimum int = scores[0]

	for _, score := range scores {

		if minimum > score {
			minimum = score
		}
	}

	return minimum
}

func getMaximum(scores ...int) int {

	var maximum int = scores[0]

	for _, score := range scores {

		if maximum < score {
			maximum = score
		}
	}

	return maximum
}

func getAvarage(scores ...int) int {

	var scoreAvarage float64

	for _, score := range scores {
		scoreAvarage += float64(score)
	}

	return int(math.Floor(scoreAvarage / float64(len(scores))))
}

func operacion(operation string) (func(scores ...int) int, error) {

	switch operation {
	case MINIMO:
		return getMinimum, nil
	case PROMEDIO:
		return getAvarage, nil
	case MAXIMO:
		return getMaximum, nil
	}

	return nil, errors.New("The operation passed is not a valid operation")
}

func main() {

	var valorMinimo int
	var valorPromedio int
	var valorMaximo int

	minFunc, err := operacion(MINIMO)
	if err != nil {
		fmt.Println(err)
	} else {
		valorMinimo = minFunc(2, 3, 4, 5, 6)

	}
	promFunc, err := operacion(PROMEDIO)
	if err != nil {
		fmt.Println(err)
	} else {
		valorPromedio = promFunc(2, 3, 4, 5, 6)

	}
	maxFunc, err := operacion(MAXIMO)
	if err != nil {
		fmt.Println(err)
	} else {
		valorMaximo = maxFunc(2, 3, 4, 5, 6)

	}

	fmt.Printf("Valor Minimo: %d \n Valor Promedio: %d \n Valor Maximo: %d \n", valorMinimo, valorPromedio, valorMaximo)
}
