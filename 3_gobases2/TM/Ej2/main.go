package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio1, error1 := promedio(10, 8, 9, 9)
	promedio2, error2 := promedio(5)
	promedio3, error3 := promedio(10, 9, -2, 4)
	fmt.Println(promedio1, error1)
	fmt.Println(promedio2, error2)
	fmt.Println(promedio3, error3)
	fmt.Println(promedio())
}

func promedio(notas ...float64) (float64, error) {
	promedio := 0.0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("error: ha ingresado una nota con valor negativo")
		}
		promedio += nota
	}
	if len(notas) > 0 {
		promedio /= float64(len(notas))
	}
	return promedio, nil
}
