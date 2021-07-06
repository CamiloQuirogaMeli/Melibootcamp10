package main

import (
	"errors"
	"fmt"
	"strings"
)

/* Para ellos necesitamos:
- Implementar una función Animal que reciba como parámetro un valor de tipo texto
con el animal especificado y nos retorne una función y un error (en caso que no exista
el animal)
- Una función para cada animal que calcule la cantidad de alimento en base a la
cantidad del tipo de animal especificado. */

func getAnimal(animal string) (func(cant int) int, error) {

	switch animal {

	case DOG:

		return getFoodForDog, nil

	case CAT:
		return getFoodForCat, nil

	case TARANTULAR:
		return getFoodForTarantula, nil

	case HAMSTER:
		return getFoodForHamster, nil

	default:
		return nil, errors.New("No contamos con alimento para tu mascota")
	}
}

func getFoodForDog(cant int) int {
	if cant > 0 {

		cant = cant * 10

	}
	return cant
}

func getFoodForCat(cant int) int {
	if cant > 0 {
		cant = cant * 5

	}
	return cant
}

func getFoodForHamster(cant int) int {
	if cant > 0 {
		cant = cant * 250
	}
	return cant
}

func getFoodForTarantula(cant int) int {
	if cant > 0 {
		var cant int
		cant = cant * 150

	}
	return cant
}

const (
	DOG        = "perro"
	CAT        = "gato"
	TARANTULAR = "tarantula"
	HAMSTER    = "hamster"
)

func main() {

	var animal string
	var cant int

	fmt.Printf("Ingrese el animal: \n")
	fmt.Scanf("%v", &animal)
	animal = strings.ToLower(animal)

	fmt.Printf("Cuantos %v tiene usted: \n", animal)
	fmt.Scanf("%v", &cant)

	animalFunc, err := getAnimal(animal)
	foodForAnimal := animalFunc(cant)

	if err == nil {
		fmt.Printf("Su %v necesita %v kg de alimento", animal, foodForAnimal)
	} else {
		fmt.Printf("Ha ocurrido un error, por favor ingrese los datos nuevamente")
	}

}
