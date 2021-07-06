/*
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear matrices en
Go.
Para ello requieren una estructura Matrix que tenga los métodos:
- Set: Recibe una matriz de flotantes e inicializa los valores en la estructura Matrix
- Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea
entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la
dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

package main

import (
	"fmt"
	"os"
)

type Matrix struct {
	Numeros []int
	Alto    int
	Ancho   int
}

func (m *Matrix) set(numeros []int) {
	m.Numeros = numeros
}

func (m *Matrix) print() {
	println("La matriz es la siguiente")
	for i, n := range m.Numeros {
		fmt.Print("  ", n, "  ")
		if i+1 == m.Ancho {
			fmt.Print("\n")
		}
	}
	fmt.Println()
	fmt.Println("El alto es: ", m.Alto, " El ancho es: ", m.Ancho)
	if m.Alto == m.Ancho {
		fmt.Println("La matriz es cuadrática")
	} else {
		fmt.Println("La matriz no es cuadrática")
	}
}

func main() {
	fmt.Println("Bienvenido al programa para crear matrices")
	fmt.Println("¿Cuanto es el alto de la matriz?")
	var alto, ancho int
	_, errAlto := fmt.Scanf("%d", &alto)
	if errAlto != nil {
		fmt.Println("Bad Input")
		os.Exit(0)
	}
	fmt.Println("¿Cuanto es el ancho de la matriz?")
	_, errAncho := fmt.Scanf("%d", &ancho)

	if errAncho != nil {
		fmt.Println("Bad Input")
		os.Exit(0)
	}
	cantidadNumeros := alto * ancho
	println(cantidadNumeros)
	numbers := []int{}
	for i := 1; i <= cantidadNumeros; i++ {
		println("Digite el ", i, "numero")
		var n int
		_, errN := fmt.Scanf("%d", &n)

		if errN != nil {
			fmt.Println("Bad Input")
			os.Exit(0)
		}
		numbers = append(numbers, n)

	}

	m := Matrix{}
	m.Alto = alto
	m.Ancho = ancho
	m.set(numbers)
	m.print()
}
