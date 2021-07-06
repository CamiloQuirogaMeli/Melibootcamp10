package main

import (
	"fmt"
)

func imprimirResultados(promedio float64, err error) {
	if err == nil {
		fmt.Println("El promedio es: ", promedio)
	} else {
		fmt.Println(err)
	}
}
