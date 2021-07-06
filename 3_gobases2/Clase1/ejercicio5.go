package main

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	tarantula = "tarantula"
	hamster   = "hamster"
)

func main() {
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
	var nombreAnimal string
	var numeroAnimales int

	fmt.Println("Digita el animal")
	fmt.Scanln(&nombreAnimal)

	funcionAnimal, err := Animal(nombreAnimal)
	if err != nil {
		fmt.Println("El", nombreAnimal, "no ha sido añadido aún")
	} else {
		fmt.Println("Digita el numero de", nombreAnimal, "s")
		fmt.Scanln(&numeroAnimales)

		cantidadAlimento, err := funcionAnimal(numeroAnimales)
		if err != nil {
			fmt.Println("Error en el cálculo del alimento")
		} else {
			fmt.Print("La cantidad de comida requerida es:")
			fmt.Println(cantidadAlimento, "kilogramos")

		}
	}

}

func Animal(tipoAnimal string) (func(int) (float64, error), error) {
	switch tipoAnimal {
	case perro:
		return alimentoPerro, nil
	case gato:
		return alimentoGato, nil
	case tarantula:
		return alimentoTarantula, nil
	case hamster:
		return alimentoHamster, nil
	}
	return nil, errors.New("El animal no ha sido añadido aún")
}

func alimentoPerro(numeroAnimales int) (float64, error) {
	if numeroAnimales <= 0 {
		return 0, errors.New("No hay animales")
	}

	cantidadAlimento := 10 * numeroAnimales

	return float64(cantidadAlimento), nil
}

func alimentoGato(numeroAnimales int) (float64, error) {
	if numeroAnimales <= 0 {
		return 0, errors.New("No hay animales")
	}

	cantidadAlimento := 5 * numeroAnimales

	return float64(cantidadAlimento), nil
}

func alimentoTarantula(numeroAnimales int) (float64, error) {
	if numeroAnimales <= 0 {
		return 0, errors.New("No hay animales")
	}
	var cantidadAlimento float64
	cantidadAlimento = (0.25 * float64(numeroAnimales))

	return float64(cantidadAlimento), nil
}

func alimentoHamster(numeroAnimales int) (float64, error) {
	if numeroAnimales <= 0 {
		return 0, errors.New("No hay animales")
	}

	var cantidadAlimento float64
	cantidadAlimento = (0.15 * float64(numeroAnimales))
	return float64(cantidadAlimento), nil
}
