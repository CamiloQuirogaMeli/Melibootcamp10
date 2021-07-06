package main

import (
	"errors"
	"fmt"
)

func main() {
	var qualifications []int
	qualification := 1

	// Se va a permitir ingresar notas hasta que se ingrese 0
	for qualification != 0 {
		fmt.Println("Ingrese una calificacion: ")
		fmt.Scanln(&qualification)
		if qualification != 0 {
			qualifications = append(qualifications, qualification)
		}
	}
	average, err := calcProm(qualifications)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio del alumno es ", average)
	}
}

func calcProm(qualifications []int) (float64, error) {
	var total int
	var average float64
	for _, qualification := range qualifications {
		if qualification < 0 {
			return 0, errors.New("Una de las calificaciones ingresadas es incorrecta")
		}
		total += qualification
	}
	average = float64(total / len(qualifications))

	return average, nil
}
