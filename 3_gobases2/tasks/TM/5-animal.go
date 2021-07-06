package main

import "errors"

func Animal(nombre string) (func(amount int) (int, error), error) {
	switch nombre {
	case "perro":
		return foodDog, nil
	case "gato":
		return foodCat, nil
	case "tarantula":
		return foodTarantula, nil
	case "hamster":
		return foodHamster, nil
	}

	notFoundError := func(amount int) (int, error) {
		return 0, nil
	}

	return notFoundError, errors.New("error: animal no encontrado")
}
