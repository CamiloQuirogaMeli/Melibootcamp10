package ejercicios

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func calcularAlimentoPerro(cantidadAnimales int) float64 {
	return float64(10 * cantidadAnimales)
}

func calcularAlimentoGato(cantidadAnimales int) float64 {
	return float64(5 * cantidadAnimales)
}

func calcularAlimentoHamster(cantidadAnimales int) float64 {
	return 0.25 * float64(cantidadAnimales)
}

func calcularAlimentoTarantula(cantidadAnimales int) float64 {
	return 0.15 * float64(cantidadAnimales)
}

func Animal(animal string) (func(cantidadAnimales int) float64, error) {
	switch animal {
	case perro:
		return calcularAlimentoPerro, nil
	case gato:
		return calcularAlimentoGato, nil
	case hamster:
		return calcularAlimentoHamster, nil
	case tarantula:
		return calcularAlimentoTarantula, nil
	default:
		return nil, errors.New("función de calculo de alimento de animal indicado no existe")
	}
}

func QuintoEjercicio() {
	animalPerro, err := Animal(perro)
	if err == nil {
		animalGato, err := Animal(gato)
		if err == nil {
			animalHamster, err := Animal(hamster)
			if err == nil {
				animalTarantula, err := Animal(tarantula)
				if err == nil {
					var cantidad float64
					cantidad += animalPerro(5)
					cantidad += animalGato(8)
					cantidad += animalHamster(30)
					cantidad += animalTarantula(2)

					fmt.Printf("Se requieren %.2f Kg de comidad", cantidad)
				}
			}
		}
	}
}
