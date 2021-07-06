package main

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

func main() {
	animalPerro, errorPerro := animal(perro)
	animalgato, errorGato := animal(gato)
	animalHamster, errorHamster := animal(hamster)
	animalTarantula, errorTarantula := animal(tarantula)

	fmt.Println("Cantidad de alimento para Perros: ", animalPerro, " Error: ", errorPerro)
	fmt.Println("Cantidad de alimento para Gatos: ", animalgato, " Error: ", errorGato)
	fmt.Println("Cantidad de alimento para Hamster: ", animalHamster, " Error: ", errorHamster)
	fmt.Println("Cantidad de alimento para Tarantula: ", animalTarantula, " Error: ", errorTarantula)
}

func animal(animal string) (float64, error) {
	switch animal {
	case perro:
		cantidadPerro := orquestador(5, opPerro)
		return cantidadPerro, nil
	case gato:
		cantidadGato := orquestador(5, opGato)
		return cantidadGato, nil
	case hamster:
		cantidadHamster := orquestador(5, opHamster)
		return cantidadHamster, nil
	case tarantula:
		cantidadTarantula := orquestador(5, opTarantula)
		return cantidadTarantula, nil
	}

	return 0, errors.New("Animal no definido")
}

func orquestador(cantidad int, operacion func(cant int) float64) float64 {
	return operacion(cantidad)
}

func opPerro(cantidad int) float64 {
	var cantidadAlimento float64 = 10
	return float64(cantidad * int(cantidadAlimento))
}

func opGato(cantidad int) float64 {
	var cantidadAlimento float64 = 5
	return float64(cantidad * int(cantidadAlimento))
}

func opHamster(cantidad int) float64 {
	var cantidadAlimento float64 = 0.250
	return float64(cantidad) * cantidadAlimento
}

func opTarantula(cantidad int) float64 {
	var cantidadAlimento float64 = 0.150
	fmt.Println("algo: ", float64(cantidad)*cantidadAlimento)
	return float64(cantidad) * cantidadAlimento
}
