package main

import (
	"errors"
	"fmt"
)

func main() {
	calc, err := Ejercicio4(PROMEDIO)

	if err != nil {
		fmt.Println("Ocurrio un error:", err)
	} else {
		resultado := calc(2, 3, 4, 1, 6, 7, 9, 15)
		fmt.Println(resultado)
	}

}

const (
	MINIMO   = "minimo"
	PROMEDIO = "promedio"
	MAXIMO   = "maximo"
)

func Ejercicio4(tipoCalculo string) (func(valores ...int) float64, error) {
	fmt.Println("Ejercicio4")
	fmt.Println("El valor", tipoCalculo, "de la lista de valores es:")
	switch tipoCalculo {
	case MINIMO:
		return minimo, nil
	case PROMEDIO:
		return promedio, nil
	case MAXIMO:
		return maximo, nil
	}
	return nil, errors.New("El calculo ingresado no esta disponible")
}

func minimo(valores ...int) float64 {
	var minimoValor float64
	for i, v := range valores {
		if i == 0 {
			minimoValor = float64(v)
		}
		if float64(v) < minimoValor {
			minimoValor = float64(v)
		}
	}
	return minimoValor
}

func promedio(valores ...int) float64 {
	var sumatoriaValores float64
	var promedioValores float64
	for _, v := range valores {
		sumatoriaValores += float64(v)
		promedioValores = sumatoriaValores / float64(len(valores))
	}
	return promedioValores
}

func maximo(valores ...int) float64 {
	var maximoValor float64
	for i, v := range valores {
		if i == 0 {
			maximoValor = float64(v)
		}
		if float64(v) > maximoValor {
			maximoValor = float64(v)
		}
	}
	return maximoValor
}
