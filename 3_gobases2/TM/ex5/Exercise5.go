package main

import (
	"errors"
	"fmt"
)

const (
	DOG       = "DOG"
	CAT       = "CAT"
	HAMSTER   = "HAMSTER"
	TARANTULA = "TARANTULA"
)

func main() {

	var foodQuantity float64

	dog, err := animal(DOG)
	cat, err := animal(CAT)
	hamster, err := animal(HAMSTER)
	tarantula, err := animal(TARANTULA)

	if err != nil {
		fmt.Println("ERROR:", err)

	}

	var dogNumber, catNumber, hamsterNumber, tarantulaNumber int

	fmt.Println("Enter the dog number:")
	fmt.Scanln(&dogNumber)
	fmt.Println("Enter the cat number:")
	fmt.Scanln(&catNumber)
	fmt.Println("Enter the hamster number:")
	fmt.Scanln(&hamsterNumber)
	fmt.Println("Enter the tarantula number:")
	fmt.Scanln(&tarantulaNumber)

	foodQuantity += dog(dogNumber)
	foodQuantity += cat(catNumber)
	foodQuantity += hamster(hamsterNumber)
	foodQuantity += tarantula(tarantulaNumber)

	fmt.Println("Food quantity:", foodQuantity)
}

func animal(animal string) (func(cantidad int) float64, error) {

	switch animal {
	case DOG:
		return calcDog, nil
	case CAT:
		return calcCat, nil
	case HAMSTER:
		return calcHamster, nil
	case TARANTULA:
		return calcTarantula, nil
	default:
		return nil, errors.New("invalid option")
	}

}

func calcDog(number int) float64 {
	return float64(number * 10)
}

func calcCat(number int) float64 {
	return float64(number * 5)
}
func calcHamster(number int) float64 {
	return float64(number) * 0.25
}
func calcTarantula(number int) float64 {
	return float64(number) * 0.15
}
