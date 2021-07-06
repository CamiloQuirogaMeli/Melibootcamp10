package ejercicio2

import (
	"testing"
)

var matrix Matrix

func TestIsQuadratic(t *testing.T) {
	m := InitializeFloatMatrix(5, 5)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			m[i][j] = float64(i*j + 1)
		}
	}
	matrix.SetMatrix(m)
	if matrix.IsQuadratic == false {
		t.Error("Test Failed: {} expected, {} received", true, false)
	}
}

func TestNotIsQuadratic(t *testing.T) {
	m := InitializeFloatMatrix(5, 6)
	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			m[i][j] = float64(i*j + 1)
		}
	}
	matrix.SetMatrix(m)
	if matrix.IsQuadratic == true {
		t.Error("Test Failed: {} expected, {} received", false, true)
	}
}

func TestHeight(t *testing.T) {
	a := InitializeFloatMatrix(5, 5)
	b := InitializeFloatMatrix(8, 5)
	c := InitializeFloatMatrix(20, 5)
	var ma, mb, mc Matrix
	ma.SetMatrix(a)
	mb.SetMatrix(b)
	mc.SetMatrix(c)
	if ma.Height != 5 || mb.Height != 8 || mc.Height != 20 {
		t.Error("Test Failed: Wrong Height")
	}

}

func TestWidth(t *testing.T) {
	a := InitializeFloatMatrix(5, 5)
	b := InitializeFloatMatrix(5, 8)
	c := InitializeFloatMatrix(5, 20)
	var ma, mb, mc Matrix
	ma.SetMatrix(a)
	mb.SetMatrix(b)
	mc.SetMatrix(c)
	if ma.Width != 5 || mb.Width != 8 || mc.Width != 20 {
		t.Error("Test Failed: Wrong Width")
	}

}

func TestMaxValue(t *testing.T) {
	var matrixA, matrixB Matrix
	matrixA.SetMatrix(InitializeFloatMatrix(5, 5))
	matrixB.SetMatrix(InitializeFloatMatrix(200, 200))
	matrixAMaxExpected := 17.0
	matrixBMaxExpected := 39602.0
	if matrixA.MaxValue != matrixAMaxExpected || matrixB.MaxValue != matrixBMaxExpected {
		t.Error("Test Failed: miscalculated maximum value")
	}
}
