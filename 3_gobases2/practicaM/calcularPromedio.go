package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(notas ...int) (int, error) {
	var resultado int

	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Las notas no pueden ser negativas!")
		}
		resultado += nota
	}
	resultado /= len(notas)
	return resultado, nil
}

func main() {

	res, error := calcularPromedio(100, 90, 50, 30, 80, 0, 100, 100, 90)

	if error != nil {
		//Si hubo error
		fmt.Println(error)
	} else {
		fmt.Println("Tu promedio es:", res)
	}

}
