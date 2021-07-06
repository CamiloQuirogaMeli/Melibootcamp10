package main

import "fmt"

const (
	suma   = "+"
	resta  = "-"
	multip = "*"
	divis  = "/"
)

func main() {
	fmt.Println(operacionAritmetica(suma, 2, 3, 2, 1, 2, 3, 4, 5, 6))
}
func operacionAritmetica(operador string, valores ...float64) float64 {
	switch operador {
	case suma:
		return orquestadorOperaciones(valores, opSuma)
	case resta:
		return orquestadorOperaciones(valores, opResta)
	case multip:
		return orquestadorOperaciones(valores, opMultip)
	case divis:
		return orquestadorOperaciones(valores, opDivis)
	}
	return 0
}
func orquestadorOperaciones(valores []float64, operacion func(value1, value2 float64) float64) float64 {
	var resultado float64
	for i, valor := range valores {
		if i == 0 {
			resultado = valor
		} else {
			resultado = operacion(resultado, valor)
		}
	}
	return resultado
}
func opSuma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}
func opResta(valor1, valor2 float64) float64 {
	return valor1 - valor2
}
func opMultip(valor1, valor2 float64) float64 {
	return valor1 * valor2
}
func opDivis(valor1, valor2 float64) float64 {
	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2
}
