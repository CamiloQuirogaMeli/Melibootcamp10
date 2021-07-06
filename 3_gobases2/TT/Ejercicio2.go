package main

import (
	"fmt"
)

type Matrix struct {
	rows int
	columns int
	isCuadratic bool
	maxValue float64
	matriz [][]float64
}

func (m *Matrix) Set(matriz [][]float64) {
	m.matriz = matriz
	m.isCuadratic, m.rows, m.columns = isCuadraticAndValues(m.matriz)
	m.maxValue = m.GetMaxValue()
}

func isCuadraticAndValues(matriz [][]float64) (bool, int, int){
	rows := len(matriz)
	columns := len(matriz[0])
	var isCuadratic bool
	if rows == columns {
		isCuadratic = true
	} else {
		isCuadratic = false
	}
	return isCuadratic, rows, columns
}

func (m Matrix) Print(){
	for i := 0; i< m.rows; i++ {
		for j := 0; j< m.columns; j++ {
			fmt.Print(m.matriz[i][j], " ")
		}
		fmt.Println("")
	}
}

func (m Matrix) GetMaxValue() float64{
	var maxValue float64 = -99999999
	for i := 0; i< m.rows; i++ {
		for j := 0; j< m.columns; j++ {
			if maxValue < m.matriz[i][j] {
				maxValue = m.matriz[i][j]
			}
		}
	}
	return maxValue
}

func main() {
	matrix := Matrix{
		rows: 2,
		columns: 3,
		isCuadratic: false,
		maxValue: 4.1,
		matriz: [][]float64{
			{1.1, 0.8, 2.5},
			{4.1, 2.0, 1.2}},
	}
	matrix.Print()
	fmt.Println("===========")
	matrix.Set([][]float64{{1.1, 0.8}, {4.1, 2.0}})
	matrix.Print()
	fmt.Println(matrix.rows, matrix.columns, matrix.maxValue, matrix.isCuadratic)
}

