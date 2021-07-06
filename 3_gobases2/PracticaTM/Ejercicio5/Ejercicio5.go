// Ejercicio 5
// Un refugio necesita calcular cuanto alimento debe comprar para las mascotas, por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan haber muchos más animales que refugiar.
// Por perro necesitan 10 kg de alimento
// Por gato 5 kg
// Por Hamsters 250 gramos.
// Por Tarántula 150 gramos.
// 3
// Para ellos necesitamos:
// - Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y nos retorne una función y un error (en caso que no exista el animal)
// - Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
// ejemplo:
// const (
// perro = "perro"
// gato = "gato"
// )
// ...
// animalPerro, err := Animal(perro)
// animalGato, err := Animal(gato)
// ...
// var cantidad float64
// cantidad += animalPerro(5)
// cantidad += animalGato(8)

package main

import (
	e "errors"
	f "fmt"
	s "strings"
)

const (
	DOG            = "dog"
	DOGALIMENT     = 10
	CAT            = "cat"
	CATALIMENT     = 5
	HAMSTER        = "hamster"
	HAMSTERALIMENT = 0.25
	SPIDERT        = "tarantula"
	SPIDERTALIMENT = 0.15
)

func calculateDogAliment(cantAnimal int) float64 {
	return float64(DOGALIMENT * cantAnimal)
}

func calculateCatAliment(cantAnimal int) float64 {
	return float64(CATALIMENT * cantAnimal)
}

func calculateHamsterAliment(cantAnimal int) float64 {
	return HAMSTERALIMENT * float64(cantAnimal)
}

func calculateSpidertAliment(cantAnimal int) float64 {
	return SPIDERTALIMENT * float64(cantAnimal)
}

func animal(animalSpecified string) (func(cantAnimal int) float64, error) {
	switch animalSpecified {
	case DOG:
		return calculateDogAliment, nil
	case CAT:
		return calculateCatAliment, nil
	case HAMSTER:
		return calculateHamsterAliment, nil
	case SPIDERT:
		return calculateSpidertAliment, nil
	}

	return nil, e.New("El animal no se encuentra entre los refugiados")
}

func verifiedAnimalSet(animals []string, animalSet string) bool {
	for _, a := range animals {
		if s.Compare(a, animalSet) == 0 {
			return true
		}
	}
	return false
}

func main() {
	var refugeeAnimals = []string{DOG, CAT, HAMSTER, SPIDERT}
	var animalsSets []string
	var mapAnimalsAliment = map[string]float64{}
	var cantAnimalSpecies int = 0
	var animalSpecie string
	var cantAnimalSpecie int = 0

	f.Printf("Los animales que se encuentran en el refugio son: ")
	for _, animal := range refugeeAnimals {
		f.Printf("%s ", animal)
	}
	f.Println()
	f.Printf("Ingresar a cuantas especies distintas (Entre los animales refugiados) les comprará alimento: ")
	f.Scan(&cantAnimalSpecies)

	for cantAnimalSpecies < 0 || cantAnimalSpecies > len(refugeeAnimals) {
		f.Printf("La cantidad total de animales es erronea\n Ingrese nuevamente: ")
		f.Scan(&cantAnimalSpecies)
	}

	for i := 0; i < cantAnimalSpecies; i++ {
		f.Printf("Ingresar la especie %d : ", i+1)
		f.Scan(&animalSpecie)
		animalSpecie = s.ToLower(animalSpecie)
		function, err := animal(animalSpecie)

		for verifiedAnimalSet(animalsSets, animalSpecie) == true || verifiedAnimalSet(refugeeAnimals, animalSpecie) == false || err != nil {
			if err != nil {
				f.Println(err)
				f.Printf("Ingrese otra especie: ")
			} else {
				f.Printf("La especie que ingresó ya se encuentra o es erronea\n Ingrese otra: ")
			}

			f.Scan(&animalSpecie)
			animalSpecie = s.ToLower(animalSpecie)
			function, err = animal(animalSpecie)
		}

		f.Printf("Indicar la cantidad de animales de la especie %s: ", animalSpecie)
		f.Scan(&cantAnimalSpecie)
		for cantAnimalSpecie < 0 {
			f.Printf("La cantidad de animales no puede ser negativa\n Ingrese nuevamente: ")
			f.Scan(&cantAnimalSpecie)
		}

		mapAnimalsAliment[animalSpecie] = function(cantAnimalSpecie)
		animalsSets = append(animalsSets, animalSpecie)
	}
	f.Printf("La comida que debe comprar para cada animal (en KG) es:\n")
	for key, element := range mapAnimalsAliment {
		f.Printf("%s : %.2f\n", key, element)
	}
}
