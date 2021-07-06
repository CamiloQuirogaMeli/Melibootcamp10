package main

import (
	"errors"
	"fmt"
)

func main() {

	res, err := average(5, 2, 4, 7, 9, 6, 3, 1, 4)
	if err != nil {
		fmt.Printf("Hay un numero negativo")
	} else {
		fmt.Printf("El promedio de las notas ingresadas es %d", res)
	}
}

func average(values ...int) (int, error) {
	var av int = 0
	for _, value := range values {
		if value >= 0 {
			av += value
		} else {
			return 0, errors.New("Hay un numero negativo")
		}

	}
	return int(av) / len(values), nil
}


