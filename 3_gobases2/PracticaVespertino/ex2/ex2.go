package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type matrix struct {
	data [][]float64
}

func main() {

	startRandSeed()

	width := 4
	height := 3

	myMatrix := inicialize(height, width)

	if !myMatrix.isEmpty() {
		myMatrix.randomFill()
		myMatrix.print()
		fmt.Println(myMatrix.getHeight())
		fmt.Println(myMatrix.getWidth())
		fmt.Println(myMatrix.getMaxValue())
		fmt.Println(myMatrix.isQuadratic())
	} else {
		fmt.Println("Warning: Empty matrix")
	}
}

func startRandSeed() {
	rand.Seed(time.Now().UnixNano())
}

func inicialize(height int, width int) (m matrix) {

	m.data = make([][]float64, height)

	for index := range m.data {
		m.data[index] = make([]float64, width)
	}

	return
}

func getRandomFloat() float64 {
	return rand.Float64()
}

func (m *matrix) randomFill() {

	for indexX := range m.data {
		for indexY := range m.data[indexX] {
			m.data[indexX][indexY] = getRandomFloat()
		}
	}
}

func (m matrix) getWidth() (height int) {

	result := 0

	for indexX := range m.data {
		result = len(m.data[indexX])
		break
	}

	return result
}

func (m matrix) getHeight() (width int) {
	return len(m.data)
}

func (m matrix) print() {

	for _, valueY := range m.data {

		for _, valueX := range valueY {
			fmt.Printf("%v\t", valueX)
		}

		fmt.Println()
	}
}

func (m matrix) getMaxValue() float64 {

	maxValue := math.MaxFloat64 * -1

	for _, valueY := range m.data {

		for _, valueX := range valueY {
			if valueX > maxValue {
				maxValue = valueX
			}
		}
	}

	return maxValue
}

func (m matrix) isQuadratic() bool {
	height := m.getHeight()
	width := m.getWidth()
	return height == width
}

func (m matrix) isEmpty() bool {
	return len(m.data) == 0
}

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
