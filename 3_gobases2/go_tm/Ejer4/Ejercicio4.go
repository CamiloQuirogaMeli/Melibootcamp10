/*
Una universidad de Entre Ríos necesita sacar estadísticas de los alumnos de cada curso,
requiriendo calcular el mínimo, máximo y promedio.
Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo,
máximo o promedio) y nos devuelve otra función ( y un error en caso que el cálculo no esté
definido) que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que
indicamos en la función anterior
*/

package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	min = "minimo"
	avg = "promedio"
	max = "maximo"
)

func main() {
	const options = `
	1. Minimo
	2. Promedio
	3. Maximo
	4. Salir
	`
	continous := true
	for continous {
		println("Que opcion desea realizar")
		println(options)
		var option int
		_, errOption := fmt.Scanf("%d", &option)

		if errOption != nil {
			fmt.Println("Bad Input")
			os.Exit(0)
		}

		switch option {
		case 1:
			notes := solicitarNotas()
			minFunc := operacion(min)
			valorMinimo, err := minFunc(notes)
			if err != nil {
				println(err)
			}
			println("La nota minima es: ", strconv.FormatFloat(valorMinimo, 'f', -2, 64))
		case 2:
			notes := solicitarNotas()
			promFunc := operacion(avg)
			valorPromedio, err := promFunc(notes)
			if err != nil {
				println(err)
			}
			println("El promedio del estudiante es: ", strconv.FormatFloat(valorPromedio, 'f', -2, 64))
		case 3:
			notes := solicitarNotas()
			maxFunc := operacion(max)
			valorMaximo, err := maxFunc(notes)
			if err != nil {
				println(err)
			}
			println("La nota maxima es: ", strconv.FormatFloat(valorMaximo, 'f', -2, 64))
		case 4:
			println("Gracias por usar el programa")
			os.Exit(0)
		default:
			println("opcion no valida")
		}
	}

}

func solicitarNotas() []float64 {
	println("¿Cuantas notas desea ingresar?")
	var quantity int
	_, errQuantity := fmt.Scanf("%d", &quantity)

	if errQuantity != nil {
		fmt.Println("Bad Input")
		os.Exit(0)
	}
	var notes = []float64{}

	for i := 1; i <= quantity; i++ {
		var note int
		println("Por favor ingrese la nota nº", i, "con decimales, Por ejemplo 15.0")

		_, errNote := fmt.Scanf("%d", &note)

		if errNote != nil {
			fmt.Println("Bad Input")
			os.Exit(0)
		}
		notes = append(notes, float64(note))

	}
	return notes

}

func operacion(opcion string) func(notas []float64) (float64, error) {
	switch opcion {
	case min:
		return opMin
	case avg:
		return opAvg
	case max:
		return opMax
	}

	return nil
}

func opMin(notas []float64) (float64, error) {
	var min float64 = notas[0]
	for _, value := range notas {
		if math.Signbit(value) {
			return 0, errors.New("Existe una nota negativa")
		}
		if value < min {
			min = value
		}

	}
	return min, nil
}
func opAvg(notas []float64) (result float64, err error) {
	for _, value := range notas {
		if math.Signbit(value) {
			return 0, errors.New("Existe una nota negativa")
		}
		result += value
	}
	return result / float64(len(notas)), nil
}
func opMax(notas []float64) (result float64, err error) {
	var max float64 = notas[0]
	for _, value := range notas {
		if math.Signbit(value) {
			return 0, errors.New("Existe una nota negativa")
		}
		if value > max {
			max = value
		}

	}
	return max, nil
}
