package main

/*Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo,
máximo o promedio) y nos devuelve otra función ( y un error en caso que el cálculo no esté
definido) que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que
indicamos en la función anterior*/
import (
	"errors"
	"fmt"
	"strings"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {
	var oper string

	fmt.Println("Ingrese que tipo de operacion desea realizar: minimo/maximo/promedio")
	fmt.Scanln(&oper)
	oper = strings.ToLower(oper)
	resultado, err := orquestador(oper, 2, 3, 3, 4, 1, 2, 4, 5)
	if err == nil {
		fmt.Println("El ", oper, " es: ", resultado)
	} else {
		fmt.Println(err)
	}

}

func orquestador(operador string, valores ...float64) (float64, error) {
	switch operador {
	case minimo:
		return minFunc(valores), nil
	case maximo:
		return maxFunc(valores), nil
	case promedio:
		return promFunc(valores), nil
	}
	return 0, errors.New("No existe ese tipo de operacion")
}

func minFunc(valores []float64) float64 {
	var minimo float64 = valores[0]

	for _, value := range valores {
		if value < minimo {
			minimo = value
		}

	}
	return minimo
}

func maxFunc(valores []float64) float64 {
	var maximo float64 = valores[0]

	for _, value := range valores {
		if value > maximo {
			maximo = value
		}

	}
	return maximo

}

func promFunc(valores []float64) float64 {
	var promedio float64 = valores[0]
	var cont int = 0
	for _, value := range valores {
		promedio = promedio + value
		cont++
	}
	if cont != 0 {
		promedio = promedio / float64(cont)
	}
	return promedio
}
