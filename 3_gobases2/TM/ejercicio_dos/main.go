package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(notas ...float32) (retorno float32, err error) {

	var suma float32
	for _, nota := range notas {

		if nota < 0 {
			retorno = 0
			err = errors.New("no puede haber una nota negativa")
			return
		}

		suma += nota
	}

	retorno = suma / float32(len(notas))
	err = nil
	return
}

func main() {

	promedio, err := calcularPromedio(4.0, 3.0, 4.0, 5.0)

	if err == nil {
		fmt.Println("promedio: ", promedio)
	} else {
		fmt.Println("Error: ", err)
	}

}
