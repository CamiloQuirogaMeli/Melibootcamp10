package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	PERRO     = "perro"
	GATO      = "gato"
	HAMSTER   = "hamster"
	TARANTULA = "tarantula"
)

func main() {
	var cantPerros, cantGatos, cantHamster, cantTarantula int
	fmt.Print("Cantidad de perros: ")
	fmt.Scanln(&cantPerros)
	fmt.Print("Cantidad de gatos: ")
	fmt.Scanln(&cantGatos)
	fmt.Print("Cantidad de hamsters: ")
	fmt.Scanln(&cantHamster)
	fmt.Print("Cantidad de tarantulas: ")
	fmt.Scanln(&cantTarantula)
	if cantPerros < 0 || cantGatos < 0 || cantHamster < 0 || cantTarantula < 0 {
		fmt.Println("Uno de los valores es inválido")
		os.Exit(1)
	}

	funcionPerro, err := Animal(PERRO)
	funcionGato, err := Animal(GATO)
	funcionHamster, err := Animal(HAMSTER)
	funcionTarantula, err := Animal(TARANTULA)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var cantidadAlimento int
	cantidadAlimento += funcionPerro(cantPerros)
	cantidadAlimento += funcionGato(cantGatos)
	cantidadAlimento += funcionHamster(cantHamster)
	cantidadAlimento += funcionTarantula(cantTarantula)

	fmt.Println("Debe comprar ", cantidadAlimento, " gramos de alimento")
}

func Animal(animal string) (func(cantidad int) int, error) {
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
		return nil, errors.New("Animal incorrecto")
	}
}
func comidaPerro(cantidad int) int {
	return cantidad * 10000
}
func comidaGato(cantidad int) int {
	return cantidad * 5000
}
func comidaHamster(cantidad int) int {
	return cantidad * 250
}
func comidaTarantula(cantidad int) int {
	return cantidad * 150
}
