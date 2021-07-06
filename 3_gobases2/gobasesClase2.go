package main

import (
	"fmt"
)

const (
	SUMA   = "+"
	RESTA  = "-"
	MULTIP = "*"
	DIVIS  = "/"
)

func bases() {
	/*
		oper := operacionAritmetica("-") //referencia a una función
		r := oper(2, 5)
		fmt.Println(r)

		s := sumatoria(2, 3, 4, 5, 6, 7, 8)
		fmt.Println(s)
	*/
	fmt.Println(operacionAritmetica(SUMA, 3, 4, 5, 6, 7, 1, 2, 3))

	s, r, m, d := operaciones(6, 2)

	fmt.Println("Suma: ", s)
	fmt.Println("Resta: ", r)
	fmt.Println("Multiplicación: ", m)
	fmt.Println("División: ", d)

	/*
		res, err := division(2, 0)
		if err != nil {
			fmt.Println(res, err)
		} else {
			fmt.Println(res)
		}
	*/

}

/*
func operacionAritmetica(operador string) func(valor1, valor2 float64) float64 {
	switch operador {
	case SUMA:
		return opSuma
	case RESTA:
		return opResta
	case MULTIP:
		return opMultip
	case DIVIS:
		return opDivis
	}
	return nil
}
*/
func operacionAritmetica(operador string, valores ...float64) float64 {
	switch operador {
	case SUMA:
		return orquestadorOperaciones(valores, opSuma)
	case RESTA:
		return orquestadorOperaciones(valores, opResta)
	case MULTIP:
		return orquestadorOperaciones(valores, opMultip)
	case DIVIS:
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
	if valor2 != 0 {
		return valor1 / valor2
	}
	return 0
}

/*
func sumatoria(values ...float64) float64 { //cantidad variable de parámetros
	var resultado float64
	for _, value := range values { //range itera los valores 1 a 1 de los parámetros que le vamos pasando y los guardamos en value
		resultado += value
	}
	return resultado
}
*/

func operaciones(valor1, valor2 float64) (suma float64, resta float64, multip float64, divis float64) { //retorno de valores nombrados
	suma = valor1 + valor2
	resta = valor1 - valor2
	multip = valor1 * valor2

	if valor2 != 0 {
		divis = valor1 / valor2
	}
	return
}

/*
func division(dividiendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("El divisor no puede ser cero")
	}
	return dividiendo / divisor, nil
}
*/
