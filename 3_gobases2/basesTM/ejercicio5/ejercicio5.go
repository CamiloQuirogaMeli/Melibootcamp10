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

func alimentoPerro(cantidad int) (int, error) {
	if cantidad < 0 {
		return 0, errors.New("no es posible procesar una cantidad negativa de perros")
	}
	return cantidad * 10000, nil
}

func alimentoGato(cantidad int) (int, error) {
	if cantidad < 0 {
		return 0, errors.New("no es posible procesar una cantidad negativa de gatos")
	}
	return cantidad * 5000, nil
}

func alimentoTarantula(cantidad int) (int, error) {
	if cantidad < 0 {
		return 0, errors.New("no es posible procesar una cantidad negativa de tarantulas")
	}
	return cantidad * 250, nil
}

func alimentoHamster(cantidad int) (int, error) {
	if cantidad < 0 {
		return 0, errors.New("no es posible procesar una cantidad negativa de hamsters")
	}
	return cantidad * 150, nil
}

func tipoAnimal(tipo string) (func(int) (int, error), error) {
	switch tipo {
	case perro:
		{
			return alimentoPerro, nil
		}
	case gato:
		{
			return alimentoGato, nil
		}
	case tarantula:
		{
			return alimentoTarantula, nil
		}
	case hamster:
		{
			return alimentoHamster, nil
		}
	default:
		{
			return nil, errors.New("el tipo de animal indicado no existe en el sistema")
		}
	}
}

func main() {
	var (
		tipo     string
		cantidad int
	)
	fmt.Println("Ingrese el tipo de animal (perro, gato, tarantula, hamster)")
	fmt.Scanln(&tipo)
	ope, err := tipoAnimal(tipo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ingrese la cantidad de animales")
		fmt.Scanln(&cantidad)
		res, errores := ope(cantidad)
		if errores != nil {
			fmt.Println(errores)
		} else {
			fmt.Println("La cantidad de alimento necesaria son ", res, " gramos")
		}
	}
}
