package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func minFunc(num ...float64) (min float64) {
	for _, values := range num {
		min = num[0]
		if values < min {
			min = values
		}
	}
	return
}

func promFunc(num ...float64) (prom float64) {
	var cant, sum float64
	for _, values := range num {
		cant++
		sum = sum + values
	}
	prom = float64(sum) / float64(cant)
	return
}

func maxFunc(num ...float64) (max float64) {
	for _, values := range num {
		max = num[0]
		if values > max {
			max = values
		}
	}
	return
}

func operacion(op string) (func(num ...float64) float64, error) {
	switch op {
	case minimo:
		return minFunc, nil
	case promedio:
		return promFunc, nil
	case maximo:
		return maxFunc, nil
	}
	return nil, errors.New("Operacion incorrecta")
}

func main() {
	minFunc, err := operacion(minimo)
	if err != nil {
		fmt.Print(err)
	} else {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Print("Valor Minimo: ", valorMinimo, "\n")
	}

	promFunc, err := operacion(promedio)
	if err != nil {
		fmt.Print(err)
	} else {
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Print("Valor Promedio: ", valorPromedio, "\n")
	}

	maxFunc, err := operacion(maximo)
	if err != nil {
		fmt.Print(err)
	} else {
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Print("Valor Maximo: ", valorMaximo, "\n")
	}

}
