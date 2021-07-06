package main

import (
	"errors"
	"fmt"
)

func calculoSalario(minutos int, categoria string) (float64, error) {
	var horas float64 = float64(minutos) / 60
	switch categoria {
	case "A":
		return (1500 * horas) * 1.5, nil
	case "B":
		return (1500 * horas) * 1.2, nil
	case "C":
		return (1000 * horas), nil
	}
	return 1, errors.New("Categoria incorrecta")
}

func main() {
	s, err := calculoSalario(60, "C")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salarios es de: ", s)
	}

}
