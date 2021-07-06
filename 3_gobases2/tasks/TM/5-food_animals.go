package main

import "errors"

func foodDog(animalAmount int) (totalFood int, err error) {
	if animalAmount < 0 {
		return 0, errors.New("error: la cantidad de perros debe ser mayor o igual a uno")
	}

	totalFood = animalAmount * 10000

	return totalFood, nil
}

func foodCat(animalAmount int) (totalFood int, err error) {
	if animalAmount < 0 {
		return 0, errors.New("error: la cantidad de gatos debe ser mayor o igual a uno")
	}

	totalFood = animalAmount * 5000

	return totalFood, nil
}

func foodHamster(animalAmount int) (totalFood int, err error) {
	if animalAmount < 0 {
		return 0, errors.New("error: la cantidad de hámster debe ser mayor o igual a uno")
	}

	totalFood = animalAmount * 250

	return totalFood, nil
}

func foodTarantula(animalAmount int) (totalFood int, err error) {
	if animalAmount < 0 {
		return 0, errors.New("error: la cantidad de hámster debe ser mayor o igual a uno")
	}

	totalFood = animalAmount * 150

	return totalFood, nil
}
