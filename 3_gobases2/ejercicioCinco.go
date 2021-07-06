package main

import (
	"errors"
	"fmt"
)

const (
	perro             = "perro"
	gato              = "gato"
	hamster           = "hamster"
	tarantula         = "tarantula"
	kilogramoEnGramos = 1000
)

func Animal(animal string) (func(num float64) float64, error) {
	switch animal {
	case perro:
		return animalPerro, nil
	case gato:
		return animalGato, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	}
	return nil, errors.New("Ocurrio un error procesando el animal")
}

func animalPerro(cantidadAnimales float64) float64 {
	var kilogramosDeAlimento float64 = 1000 / kilogramoEnGramos
	return kilogramosDeAlimento * cantidadAnimales
}

func animalGato(cantidadAnimales float64) float64 {
	var kilogramosDeAlimento float64 = 5000 / kilogramoEnGramos
	return kilogramosDeAlimento * cantidadAnimales
}

func animalHamster(cantidadAnimales float64) float64 {
	var kilogramosDeAlimento float64 = 250 / kilogramoEnGramos
	return kilogramosDeAlimento * cantidadAnimales
}

func animalTarantula(cantidadAnimales float64) float64 {
	var kilogramosDeAlimento float64 = 150 / kilogramoEnGramos
	return kilogramosDeAlimento * cantidadAnimales
}

func main() {
	animalPerro, err := Animal(perro)
	animalGato, err := Animal(gato)

	if err != nil {
		fmt.Println("Ocurrio un error")
	}

	var cantidad float64
	cantidad += animalPerro(8)
	cantidad += animalGato(8)
	cantidad += animalHamster(4)
	cantidad += animalTarantula(12)

	fmt.Printf("La cantidad de alimento necesaria es de %.2f Kilogramos\n", cantidad)
}
