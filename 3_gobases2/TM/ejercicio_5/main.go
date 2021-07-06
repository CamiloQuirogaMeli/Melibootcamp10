package main

import (
	"fmt"
	"errors"
)

const (
	PERRO = "perro"
	GATO = "gato"
	HAMSTER = "hamster"
	TARANTULA = "tarantula"
)

func foodGramsForHamsters(amount int) int {
	return amount * 250
}

func foodGramsForDogs(amount int) int {
	return amount * 10000
}

func foodGramsForCats(amount int) int {
	return amount * 5000
}

func foodGramsForTarantulas(amount int) int {
	return amount * 150
}

func Animal(animalType string) (func(foodAmount int) int, error) {
	switch animalType {
		case HAMSTER: 
			return foodGramsForHamsters, nil
		case PERRO:	
			return foodGramsForDogs, nil
		case GATO:	
			return foodGramsForCats, nil
		case TARANTULA:	
			return foodGramsForTarantulas, nil
	}
	return nil, errors.New("El tipo de animal no existe")
}

func main() {
	var cantidad int

	fmt.Println("* Caso comida para hamster")
	animalHamster, _ := Animal(HAMSTER)
	cantidad = animalHamster(5)
	fmt.Println("\tCantidad de hamsters:", 5)
	fmt.Printf("\tComida necesaria: %dg\n", cantidad)

	fmt.Println("* Caso comida para perro")
	animalPerro, _ := Animal(PERRO)
	cantidad = animalPerro(2)
	fmt.Println("\tCantidad de perros:", 2)
	fmt.Printf("\tComida necesaria: %dg\n", cantidad)

	fmt.Println("* Caso comida para gato")
	animalGato, _ := Animal(GATO)
	cantidad = animalGato(3)
	fmt.Println("\tCantidad de gatos:", 3)
	fmt.Printf("\tComida necesaria: %dg\n", cantidad)

	fmt.Println("* Caso comida para tarantula")
	animalTarantula, _ := Animal(TARANTULA)
	cantidad = animalTarantula(7)
	fmt.Println("\tCantidad de tarantulas:", 7)
	fmt.Printf("\tComida necesaria: %dg\n", cantidad)

	fmt.Println("* Caso animal inexistente")
	_, err := Animal("abc")
	fmt.Println("\tError:", err)
}