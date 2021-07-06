package main

import (
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
	// err       = "error"
)

func main() {
	animalPerro, errPerro := Animal(perro)
	animalGato, errGato := Animal(gato)
	animalTarántula, errTarántula := Animal(tarantula)
	animalHámster, errHámster := Animal(hamster)
	// _, errError := Animal(err)

	if errPerro != nil {
		fmt.Println(errPerro)
	}

	if errGato != nil {
		fmt.Println(errGato)
	}

	if errTarántula != nil {
		fmt.Println(errTarántula)
	}

	if errHámster != nil {
		fmt.Println(errHámster)
	}

	// if errError != nil {
	// 	fmt.Println(errError)
	// }

	var totalFood int
	var pFood int

	pFood, errPerro = animalPerro(10)
	totalFood += pFood
	pFood, errGato = animalGato(64)
	totalFood += pFood
	pFood, errTarántula = animalTarántula(124)
	totalFood += pFood
	pFood, errHámster = animalHámster(300)
	totalFood += pFood

	fmt.Println("La cantidad de comida total para el refugio es:", totalFood, "gr de alimento")
}
