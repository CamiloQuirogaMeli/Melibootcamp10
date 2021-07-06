// Ejercicio 4 - Calcular estadísticas
// Una universidad de Entre Ríos necesita sacar estadísticas de los alumnos de cada curso,
// requiriendo calcular el mínimo, máximo y promedio.
// Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo, máximo o promedio)
// y nos devuelve otra función ( y un error en caso que el cálculo no esté definido)
// que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que indicamos en la función anterior
// Ejemplo:
// const (
// minimo = "minimo"
// promedio = "promedio"
// maximo = "maximo"
// )
// ...
// minFunc, err := operacion(minimo)
// promFunc, err := operacion(promedio)
// maxFunc, err := operacion(maximo)
// ...
// valorMinimo := minFunc(2,3,3,4,1,2,4,5)
// valorPromedio := promFunc(2,3,3,4,1,2,4,5)
// valorMaximo := maxFunc(2,3,3,4,1,2,4,5)

package main

import (
	e "errors"
	f "fmt"
)

const (
	aver = "Promedio"
	min  = "Minimo"
	max  = "Maximo"
)

func averageFunc(values []float64) float64 {
	var total float64 = 0
	for _, value := range values {
		total += float64(value)
	}
	return total / float64(len(values))
}

func maximumFunc(values []float64) float64 {
	maximumAux := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > maximumAux {
			maximumAux = values[i]
		}
	}
	return float64(maximumAux)
}

func minimumFunc(values []float64) float64 {
	minimumAux := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < minimumAux {
			minimumAux = values[i]
		}
	}
	return float64(minimumAux)
}

func mathematicOperations(operation string) (func(values []float64) float64, error) {

	switch operation {
	case aver:
		return averageFunc, nil
	case min:
		return minimumFunc, nil
	case max:
		return maximumFunc, nil
	}

	return nil, e.New("No eligio una operacion valida\n")
}

func main() {

	var option string
	var cantNumbers int
	var values []float64
	var number float64

	f.Printf("Operaciones disponibles:\n-Promedio\n -Minimo\n -Maximo\n")
	f.Printf("¿Que operacion desea realizar?: ")
	f.Scan(&option)
	function, err := mathematicOperations(option)

	if err != nil {
		f.Printf("%s", err)
		return
	}

	f.Printf("Ingrese la cantidad de valores para operar: ")
	f.Scan(&cantNumbers)

	if cantNumbers < 0 {
		f.Printf("Usted no desea agregar ningún valor")
		return
	}

	for i := 0; i < cantNumbers; i++ {
		f.Printf("Ingrese el valor %d: ", i+1)
		f.Scan(&number)
		values = append(values, number)
	}

	f.Printf("Usted eligio la operacion %s y el resultado es: %.2f", option, function(values))

}
