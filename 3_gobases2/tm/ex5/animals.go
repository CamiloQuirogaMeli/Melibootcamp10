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

func animal(ani string) (func(cantidadAnimal int) float64, error) {
	switch ani {
	case PERRO:
		return animalPerro, nil
	case GATO:
		return animalGato, nil
	case HAMSTER:
		return animalHamster, nil
	case TARANTULA:
		return animalTarantula, nil
	}
	return nil, errors.New("Animal no encontrado")
}

func animalPerro(cantidad int) float64 {
	return float64(cantidad) * 10
}
func animalGato(cantidad int) float64 {
	return float64(cantidad) * 5
}
func animalHamster(cantidad int) float64 {
	return float64(cantidad) * 0.250
}
func animalTarantula(cantidad int) float64 {
	return float64(cantidad) * 0.150
}

func main() {
	var cantidad float64

	dog, err := animal(PERRO)
	cat, err := animal(GATO)
	hams, err := animal(HAMSTER)
	taran, err := animal(TARANTULA)

	if err != nil {
		fmt.Println(err)
	} else {
		cantidad = dog(2)
		fmt.Printf("La cantidad a comprar de comida para perros es de %.2fkg\n", cantidad)
		cantidad = cat(5)
		fmt.Printf("La cantidad a comprar de comida para gatos es de %.2fkg\n", cantidad)
		cantidad = hams(1)
		fmt.Printf("La cantidad a comprar de comida para hamsters es de %.2fkg\n", cantidad)
		cantidad = taran(3)
		fmt.Printf("La cantidad a comprar de comida para tarantulas es de %.2fkg\n", cantidad)
	}

}
