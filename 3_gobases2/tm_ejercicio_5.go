package main

import (
	"errors"
	"fmt"
)

func main() {

	animal := "perro"
	cantidadAnimales := 10

	animalFunc, err := Animal(animal)

	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}

	fmt.Printf("La cantidad de alimento necesaria es: %f \n", animalFunc(cantidadAnimales))

}

func Animal(animal string) (func(int) float32, error) {

	switch animal {
	case "tarantulas":
		return alimentoTarantula, nil
	case "hamster":
		return alimentoHamster, nil
	case "perro":
		return alimentoPerro, nil
	case "gato":
		return alimentoGato, nil
	default:
		return nil, errors.New("No esta definido el animal")

	}
}

func alimentoTarantula(cantidad int) float32 {
	return 0.150 * float32(cantidad)
}

func alimentoHamster(cantidad int) float32 {
	return 0.250 * float32(cantidad)
}

func alimentoPerro(cantidad int) float32 {
	return float32(10 * cantidad)
}

func alimentoGato(cantidad int) float32 {
	return float32(5 * cantidad)
}
