package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	print := fmt.Println
	animal := "tarantula"
	getFoodForAnimal, err := animalHandler(animal)
	if err == nil {
		foodAmount := getFoodForAnimal(2)
		print("We need", foodAmount, "kg for your animals")
	} else {
		print(err)
	}
}

func getFoodForDog(numDogs int) float64 {
	return float64(numDogs) * 10.0
}
func getFoodForCat(numCats int) float64 {
	return float64(numCats) * 5.0
}
func getFoodForHamster(numHamsters int) float64 {
	return float64(numHamsters) * 0.25
}
func getFoodForTarantula(numTarantulas int) float64 {
	return float64(numTarantulas) * 0.15
}

func animalHandler(animal string) (func(numAnimals int) float64, error) {
	animalFunctions := map[string]func(numAnimals int) float64{
		dog:       getFoodForDog,
		cat:       getFoodForCat,
		hamster:   getFoodForHamster,
		tarantula: getFoodForTarantula}

	if function, ok := animalFunctions[animal]; ok {
		return function, nil
	} else {
		return nil, errors.New(animal + " is not an animal we take care of here")
	}

}
