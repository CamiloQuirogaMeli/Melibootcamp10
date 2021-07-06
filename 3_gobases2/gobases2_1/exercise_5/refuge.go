package main

import (
	er "errors"
	f "fmt"
)

var result float32
var total float32

const (
	Dog       = "dog"
	Cat       = "cat"
	Hamster   = "hamster"
	Tarantula = "tarantula"
	Fish      = "fish"
)

func calculateDog(dog float32) float32 {
	result = dog * 10
	return result
}

func calculateCat(cat float32) float32 {
	result = cat * 5
	return result
}

func calculateHamster(hamster float32) float32 {
	result = hamster * 0.25
	return result
}

func calculateTarantula(tarantula float32) float32 {
	result = tarantula * 0.15
	return result
}

func main() {

	dogFuction, err1 := selectAnimal(Dog)
	handleErrors(err1)

	catFuction, err2 := selectAnimal(Cat)
	handleErrors(err2)

	_, err3 := selectAnimal(Fish)
	handleErrors(err3)

	hamsterFuction, err3 := selectAnimal(Hamster)
	handleErrors(err3)

	total += dogFuction(3)
	total += catFuction(5)
	total += hamsterFuction(8)

	f.Println("Must buy food", total, "Kg for pets")

}

func selectAnimal(animal string) (func(animal float32) float32, error) {

	switch animal {

	case Dog:
		return calculateDog, nil

	case Cat:
		return calculateCat, nil

	case Hamster:
		return calculateHamster, nil

	case Tarantula:
		return calculateTarantula, nil

	default:
		return nil, er.New("There is no record of this animal")
	}

}

func handleErrors(err error) {
	if err != nil {
		f.Println(err)
	}
}
