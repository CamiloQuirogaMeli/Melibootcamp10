package main

import (
	"errors"
	"fmt"
)

const (
	PERRO     = "perro"
	GATO      = "gato"
	HAMSTER   = "hamster"
	TARANTULA = "tarantula"
)

func main() {
	/*
		Ejercicio 5
		Un refugio necesita calcular cuanto alimento debe comprar para las mascotas, por el
		momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan haber
		muchos más animales que refugiar.

		Por perro necesitan 10 kg de alimento
		Por gato 5 kg
		Por Hamsters 250 gramos.
		Por Tarántula 150 gramos.

		Para ellos necesitamos:
		- Implementar una función Animal que reciba como parámetro un valor de tipo texto
		con el animal especificado y nos retorne una función y un error (en caso que no exista
		el animal)
		- Una función para cada animal que calcule la cantidad de alimento en base a la
		cantidad del tipo de animal especificado.
	*/

	animalPerro, err := Animal(PERRO)
	animalGato, err := Animal(GATO)
	animalHamster, err := Animal(HAMSTER)
	animalTarantula, err := Animal(TARANTULA)

	if err != nil {
		fmt.Println(err)
	} else {
		totalPerro := animalPerro(5)
		totalGato := animalGato(8)
		totalHamster := animalHamster(10)
		totalTarantula := animalTarantula(3)

		fmt.Println(
			"Comida para Perros", totalPerro, "\n",
			"Comida para Gatos", totalGato, "\n",
			"Comida para Hamsters", totalHamster, "\n",
			"Comida para Tarantulas", totalTarantula)
	}

}

func Animal(animal string) (func(val int) float64, error) {

	switch animal {
	case PERRO:
		return comidaPerro, nil
	case GATO:
		return comidaGato, nil
	case HAMSTER:
		return comidaHamster, nil
	case TARANTULA:
		return comidaTarantula, nil
	default:
		return nil, errors.New("No existe animal")
	}
}

func comidaPerro(cant int) float64 {
	return float64(cant) * 10
}

func comidaGato(cant int) float64 {
	return float64(cant) * 5
}

func comidaHamster(cant int) float64 {
	return float64(cant) * 0.25
}

func comidaTarantula(cant int) float64 {
	return float64(cant) * 0.15
}
