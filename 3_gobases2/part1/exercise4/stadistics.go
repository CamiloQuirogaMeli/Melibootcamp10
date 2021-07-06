package main

import (
	e "errors"
	f "fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func operacion(operacion string) (func(values ...int) float32, error) {
	switch operacion {
	case minimo:
		return min, nil
	case promedio:
		return prom, nil
	case maximo:
		return max, nil
	default:
		msg := "undefined " + operacion + " function"
		return nil, e.New(msg)
	}
}

func min(values ...int) float32 {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return float32(min)
}

func prom(values ...int) float32 {
	prom := values[0]
	for _, value := range values {
		prom += value
	}
	return float32(prom / len(values))
}

func max(values ...int) float32 {
	maxi := values[0]
	for _, value := range values {
		if value > maxi {
			maxi = value
		}
	}
	return float32(maxi)
}

func main() {
	var _errors []error

	minFunc, err1 := operacion(minimo)
	promFunc, err2 := operacion(promedio)
	maxFunc, err3 := operacion(maximo)

	switch {
	case err1 != nil:
		_errors = append(_errors, err1)
		fallthrough
	case err2 != nil:
		_errors = append(_errors, err2)
		fallthrough
	case err3 != nil:
		_errors = append(_errors, err3)
	default:
		break
	}

	if len(_errors) > 0 {
		for i, err := range _errors {
			f.Println("ERROR #:", i+1, err)
		}
	} else {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		f.Println("Min: ", valorMinimo)
		f.Println("Avg: ", valorPromedio)
		f.Println("Max: ", valorMaximo)
	}

}
