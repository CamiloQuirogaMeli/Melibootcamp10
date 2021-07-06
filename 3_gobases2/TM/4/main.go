package main
import (
	"errors"
	"fmt"
)

/* Ejercicio 4 - Calcular estadísticas
Una universidad de Entre Ríos necesita sacar estadísticas de los alumnos de cada curso,
requiriendo calcular el mínimo, máximo y promedio.

Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo,
máximo o promedio) y nos devuelve otra función ( y un error en caso que el cálculo no esté
definido) que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que
indicamos en la función anterior
Ejemplo:

const (
minimo = "minimo"
promedio = "promedio"
maximo = "maximo"
)

...
minFunc, err := operacion(minimo)
promFunc, err := operacion(promedio)
maxFunc, err := operacion(maximo)

...
valorMinimo := minFunc(2,3,3,4,1,2,4,5)
valorPromedio := promFunc(2,3,3,4,1,2,4,5)
valorMaximo := maxFunc(2,3,3,4,1,2,4,5) */

const (
	minimo = "minimo"
	promedio = "promedio"
	maximo = "maximo"
)

func main() {
	f, err := operacion(maximo)

	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println("El resultado de la operacion es:", f(2, 3, 4, 5, 6, 7, 8, 9, 10))
	}
}

func operacion(o string) (func(i... int) (int), error) {
	switch o {
		case minimo: {
			return func (i... int) int {
				min := 11
				for _, val := range i {
					if (val < min) {
						min =  val
					}
				}
				return min
			}, nil
		}
		case promedio: {
			return func (i... int) int {
				suma, cant := 0, 0
				for _, val := range i {
					suma += val
					cant++
				}
				return (suma / cant)
			}, nil
		}
		case maximo: {
			return func (i... int) int {
				max := 0
				for _, val := range i {
					if (val > max) {
						max =  val
					}
				}
				return max
			}, nil
		}
		default: return nil, errors.New("Valor invalido.")
	}
}