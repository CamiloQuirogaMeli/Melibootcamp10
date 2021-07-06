package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Matrix struct {
	matrix    [][]float64
	height    int
	width     int
	maxValue  float64
	isSquared bool
}

func (m *Matrix) setMatrix(theMatrix [][]float64) error {
	m.matrix = theMatrix
	m.height = len(theMatrix)
	if m.height == 0 {
		return errors.New("Matrix can't be empty")
	} else {
		m.width = len(theMatrix[0])
		m.maxValue = m.getMaxValue()
		m.isSquared = m.width == m.height
		return nil
	}
}

func (m Matrix) toString() string {
	var stringMatrix string = ""
	for row := 0; row < m.height; row++ {
		thisRow := m.matrix[row]
		for column := 0; column < m.width; column++ {
			stringMatrix += strconv.FormatFloat(thisRow[column], 'f', -1, 64)
			if column != m.width-1 {
				stringMatrix += "\t"
			}
		}
		if row != m.height-1 {
			stringMatrix += "\n"
		}
	}
	return stringMatrix
}

func (m Matrix) getMaxValue() float64 {
	max := math.Inf(-1)
	for row := 0; row < m.height; row++ {
		thisRow := m.matrix[row]
		for column := 0; column < m.width; column++ {
			max = math.Max(max, thisRow[column])
		}
	}
	return max
}

func (m Matrix) printDetails() {
	fmt.Println(m.toString())
	fmt.Println("Maximum value of this matrix is: ", m.maxValue)
	aux := ""
	if !m.isSquared {
		aux = " not"
	}
	fmt.Print("This matrix is", aux, " squared\n")

}

func main() {
	fmt.Println("First matrix example\n")
	myMatrix := Matrix{}
	err := myMatrix.setMatrix([][]float64{{1, 8.9, 5}, {3, 4, 0}, {8, 1, 3}})
	if err != nil {
		fmt.Println(err)
	} else {
		myMatrix.printDetails()
	}

	fmt.Println("\nSecond matrix example\n")
	err = myMatrix.setMatrix([][]float64{})
	if err != nil {
		fmt.Println(err)
	} else {
		myMatrix.printDetails()
	}

	fmt.Println("\nThird matrix example\n")
	err = myMatrix.setMatrix([][]float64{{1, 2}, {3, 4}, {5, 6}})
	if err != nil {
		fmt.Println(err)
	} else {
		myMatrix.printDetails()
	}

}
