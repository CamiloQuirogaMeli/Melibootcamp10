package main

import (
	"errors"
	"fmt"
)

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

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func seleccionaAnimal(animal string) (func(float64) (float64, error), error) {
	switch animal {
	case perro:
		return calculoPerro, nil
	case gato:
		return calculoGato, nil
	case hamster:
		return calculoHamster, nil
	case tarantula:
		return calculoTarantula, nil
	}
	return nil, errors.New("Animal no existe")
}

func calculoPerro(cantidad float64) (float64, error) {
	if cantidad <= 0 {
		return 0, errors.New("No se aceptan valores negativos o cero")
	}
	resultado := cantidad * 10 // Kilogramos
	return resultado, nil
}

func calculoGato(cantidad float64) (float64, error) {
	if cantidad <= 0 {
		return 0, errors.New("No se aceptan valores negativos o cero")
	}
	resultado := cantidad * 5 // Kilogramos
	return resultado, nil
}

func calculoHamster(cantidad float64) (float64, error) {
	if cantidad <= 0 {
		return 0, errors.New("No se aceptan valores negativos o cero")
	}
	resultado := cantidad * 250 // Gramos
	return resultado, nil
}

func calculoTarantula(cantidad float64) (float64, error) {
	if cantidad <= 0 {
		return 0, errors.New("No se aceptan valores negativos o cero")
	}
	resultado := cantidad * 150 // Gramos
	return resultado, nil
}

func main() {
	animal := ""
	fmt.Print("Digite el animal al que desea calcular la comida:\nperro\ngato\nhamster\ntarantula")
	fmt.Scan(&animal)

	mifuncion, err := seleccionaAnimal(animal)

	if err != nil {
		fmt.Println(err)
	} else {
		cantidad := 0.0
		fmt.Println("Digite la cantidad de animales para", animal)
		fmt.Scan(&cantidad)
		respuesta, errors := mifuncion(cantidad)
		if errors != nil {
			fmt.Println(errors)
		} else {
			fmt.Println(respuesta)
		}
	}
}
