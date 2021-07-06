package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	maximo   = "maximo"
	promedio = "promedio"
)

func calculoMinimo(numeros ...float64) (float64, error) {
	if len(numeros) == 0 {
		return 0, errors.New("no hay datos sobre los alumnos")
	}
	minimo := numeros[0]
	for _, value := range numeros {
		if value < 0 {
			return 0, errors.New("no pueden procesarse numeros negativos")
		}
		if value < minimo {
			minimo = value
		}
	}
	return minimo, nil
}

func calculoMaximo(numeros ...float64) (float64, error) {
	var maximo float64
	if len(numeros) == 0 {
		return 0, errors.New("no hay datos sobre los alumnos")
	}
	for _, value := range numeros {
		if value < 0 {
			return 0, errors.New("no pueden procesarse numeros negativos")
		}
		if value > maximo {
			maximo = value
		}
	}
	return maximo, nil
}

func calculoPromedio(numeros ...float64) (float64, error) {
	var promedio float64
	if len(numeros) == 0 {
		return 0, errors.New("no hay datos sobre los alumnos")
	}
	for _, value := range numeros {
		if value < 0 {
			return 0, errors.New("no pueden procesarse numeros negativos")
		}
		promedio += value
	}
	return promedio / float64(len(numeros)), nil
}

func manejadorOperaciones(operacion string) (func(...float64) (float64, error), error) {
	switch operacion {
	case minimo:
		{
			return calculoMinimo, nil
		}
	case maximo:
		{
			return calculoMaximo, nil
		}
	case promedio:
		{
			return calculoPromedio, nil
		}
	}
	return nil, errors.New("operacion indicada no existe")
}

func main() {
	var operacion string
	fmt.Println("Los resultados de los estudiantes son 7, 8, 5, 2")
	fmt.Println("Ingrese la operacion a realizar (maximo, minimo, promedio)")
	fmt.Scanln(&operacion)
	ope, err := manejadorOperaciones(operacion)
	if err != nil {
		fmt.Println(err)
	} else {
		res, errores := ope(7, 8, 5, 2)
		if errores != nil {
			fmt.Println(errores)
		} else {
			fmt.Println("El ", operacion, " es ", res)
		}
	}

}
