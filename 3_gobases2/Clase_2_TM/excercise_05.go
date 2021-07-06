package main

import (
	"errors"
	"fmt"
)

const (
	DOG          = "perro"
	CAT          = "gato"
	HAMSTER      = "hamster"
	SPIDER       = "spider"
	DOG_FOOD     = 10000.00
	CAT_FOOD     = 5000.00
	HAMSTER_FOOD = 250.00
	SPIDER_FOOD  = 150.00
)

func main() {

	animalDog, _ := Animal(DOG)
	animalCat, _ := Animal(CAT)

	animalLion, err := Animal("Lion")

	var quantityAnimalFood float64

	quantityAnimalFood += animalDog(2)
	quantityAnimalFood += animalCat(3)

	fmt.Printf("Animal food kg [%2.f]\n", quantityAnimalFood)

	if animalLion == nil {
		fmt.Println(err)
	}

}

func Animal(animal string) (func(int) float64, error) {
	switch animal {
	case DOG:
		return dogFood, nil
	case CAT:
		return catFood, nil
	case HAMSTER:
		return hamsterFood, nil
	case SPIDER:
		return spiderFood, nil
	}

	return nil, errors.New(fmt.Sprintf("Cannot food for animal: [%s]", animal))
}

func dogFood(quantity int) float64 {
	return float64(quantity) * DOG_FOOD
}

func catFood(quantity int) float64 {
	return float64(quantity) * CAT_FOOD
}

func hamsterFood(quantity int) float64 {
	return float64(quantity) * HAMSTER_FOOD
}

func spiderFood(quantity int) float64 {
	return float64(quantity) * SPIDER_FOOD
}
