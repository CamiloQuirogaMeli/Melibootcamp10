package main

import (
	"errors"
	"fmt"
)

const (
	PERRO     = "perro"
	GATO      = "gato"
	TARANTULA = "tarantula"
	HAMSTER   = "hamster"
)

var (
	pets    = map[string]int{PERRO: 6, GATO: 12, TARANTULA: 33, HAMSTER: 40}
	kgToBuy float64
)

func main() {

	kgToBuy = 0.0

	for pet, cant := range pets {

		//fmt.Println(pet)

		petFnc, err := animal(pet)

		if err == nil {
			kgToBuy += petFnc(cant)
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println(kgToBuy)
}

func animal(pet string) (func(int) float64, error) {

	switch pet {
	case "perro":
		return comidaPerro, nil
	case "gato":
		return comidaGato, nil
	case "hamster":
		return comidaHamster, nil
	case "tarantula":
		return comidaTarantula, nil
	}

	return nil, errors.New("pet not fund")
}

func comidaPerro(cant int) float64 {
	return (float64)(cant) * 10.0
}

func comidaGato(cant int) float64 {
	return (float64)(cant) * 5.0
}

func comidaHamster(cant int) float64 {
	return (float64)(cant) * 0.25
}

func comidaTarantula(cant int) float64 {
	return (float64)(cant) * 0.15
}

/*
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
