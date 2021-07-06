package main

import "errors"

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func comidaPorAnimal(animal string) (func(value float64) float64, error) {
	switch animal {
	case perro:
		return comidaPerro, nil
	case gato:
		return comidaGato, nil
	case hamster:
		return comidaHamster, nil
	case tarantula:
		return comidaTarantula, nil
	}
	msg := errors.New("Ese animal no existe")
	return nil, msg
}

func comidaPerro(cant float64) (cantidadComida float64) {
	var comidaP float64 = 10
	cantidadComida = comidaP * cant

	return
}

func comidaGato(cant float64) (cantidadComida float64) {
	var comidaG float64 = 5
	cantidadComida = comidaG * cant
	return
}

func comidaHamster(cant float64) (cantidadComida float64) {
	var comidaH float64 = 250
	cantidadComida = comidaH * cant
	return
}

func comidaTarantula(cant float64) (cantidadComida float64) {
	var comidaT float64 = 150
	cantidadComida = comidaT * cant
	return
}
