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

func operacion(tipoValor string) (func(num ...int) int, error) {
	/*switch tipoValor {
	case "minimo":
		return minFunc, nil
	case "promedio":
		return promFunc, nil
	case "maximo":
		return maxFunc, nil
	}
	return nil, errors.New("No se pudo calcular")*/
	if tipoValor == "minimo" {
		numeroMinimo := func(num ...int) int {
			var numMinimo int = 10000
			for _, valor := range num {
				if valor < numMinimo {
					numMinimo = valor
				}
			}
			return numMinimo
		}
		return numeroMinimo, nil
	}
	if tipoValor == "promedio" {
		numeroPromedio := func(num ...int) int {
			var promedio int
			var suma int
			for _, valor := range num {
				suma += valor
			}
			promedio = suma / len(num)
			return promedio
		}
		return numeroPromedio, nil
	}
	if tipoValor == "maximo" {
		numeroMaximo := func(num ...int) int {
			var numMaximo int
			for _, valor := range num {
				if valor > numMaximo {
					numMaximo = valor
				}
			}
			return numMaximo
		}
		return numeroMaximo, nil
	}
	return nil, errors.New("No se pudo calcular")
}

func minFunc(valor ...int) (int, error) {
	var minimo int
	flag := false
	for _, value := range valor {
		if flag == false {
			minimo = value
			flag = true
		}
		if minimo > value {
			minimo = value
		}
	}
	return minimo, nil
}
func promFunc(valor ...int) (int, error) {
	var promedio int = 0
	for _, value := range valor {
		promedio += value
	}
	return promedio / len(valor), nil
}
func maxFunc(valor ...int) (int, error) {
	var maximo int
	flag := false
	for _, value := range valor {
		if flag == false {
			maximo = value
			flag = true
		}
		if maximo > value {
			maximo = value
		}
	}
	return maximo, nil
}

func main() {
	minFunc, err := operacion(minimo)
	promFunc, err := operacion(promedio)
	maxFunc, err := operacion(maximo)

	if err != nil {
		fmt.Println("Ocurrio un error")
	}

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(valorMinimo)
	fmt.Println(valorPromedio)
	fmt.Println(valorMaximo)

}
