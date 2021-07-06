package main

import (
	"errors"
	"fmt"
)

func promedio(notas ...int) (int, error) {
	var resultado int
	for _, value := range notas {
		if value < 0 {
			return 0, errors.New("verifique los numeros ingresados")
		}
		resultado += value
	}
	return resultado / len(notas), nil
}

func main() {
	res, err := promedio(10, 3, 10, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es de:", res)
	}
}
