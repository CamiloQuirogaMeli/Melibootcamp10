package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "perro"
	cat       = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func foodForDogsGr(n int) int {
	return n * 10000
}

func foodForCatsGr(n int) int {
	return n * 5000
}

func foodForHamstersGr(n int) int {
	return n * 250
}

func foodForTarantulasGr(n int) int {
	return n * 150
}

func animal(animalType string) (func(n int) int, error) {
	animalFunctionMap := map[string]func(n int) int{}
	animalFunctionMap[dog] = foodForDogsGr
	animalFunctionMap[cat] = foodForCatsGr
	animalFunctionMap[hamster] = foodForHamstersGr
	animalFunctionMap[tarantula] = foodForTarantulasGr

	if function, ok := animalFunctionMap[animalType]; ok {
		return function, nil
	}
	return nil, errors.New("No hay datos de ese animal")
}

func main() {
	animalType := dog
	function, err := animal(animalType)
	if err == nil {
		fmt.Println("Los animales", animalType, "necesitan", function(3), "gramos de comida")
	} else {
		fmt.Println(err)
	}

	//Tests described in the excercise statement
	animalPerro, err := animal(dog)
	animalGato, err := animal(cat)
	animalHamster, err := animal(hamster)
	animalTarantula, err := animal(tarantula)

	var cantidad int
	cantidad += animalPerro(5)
	cantidad += animalGato(8)
	cantidad += animalHamster(8)
	cantidad += animalTarantula(8)

	fmt.Println("Cantidad de comida para todos los animales", cantidad, "gramos")

}
