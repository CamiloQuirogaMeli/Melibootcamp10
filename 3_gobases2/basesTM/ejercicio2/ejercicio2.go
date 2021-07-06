package main

import (
	"errors"
	"fmt"
)

func promediar(numeros ...float64) (float64, error) {
	var resultado float64
	for _, value := range numeros {
		if value < 0 {
			return 0, errors.New("no pueden procesarse numeros negativos")
		}
		resultado += value
	}
	return resultado / float64(len(numeros)), nil
}

func main() {
	var res float64
	var err error
	fmt.Println("pasando los numeros 7, 8, 3, 4")
	res, err = promediar(7, 8, 3, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio calculado es ", res)
	}
	fmt.Println("pasando los numeros 7, 8, -6, 4, 0, 2")
	res, err = promediar(7, 8, -6, 4, 0, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio calculado es ", res)
	}
}
