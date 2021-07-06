package main

import (
	"errors"
	"fmt"
	"os"
)

// food expressed in kg
const (
	foodByDog       = 10.0
	foodByCat       = 5.0
	foodByHamster   = 0.250
	foodByTarantula = 0.150
)

func amountDogFood(amount int) float64 {
	return float64(amount) * foodByDog
}

func amountCatFood(amount int) float64 {
	return float64(amount) * foodByCat
}

func amountHamsterFood(amount int) float64 {
	return float64(amount) * foodByHamster
}

func amountTarantulaFood(amount int) float64 {
	return float64(amount) * foodByTarantula
}

func animal(tipo string) (func(amount int) float64, error) {
	switch tipo {
	case "perro":
		return amountDogFood, nil
	case "gato":
		return amountCatFood, nil
	case "hamster":
		return amountHamsterFood, nil
	case "tarantula":
		return amountTarantulaFood, nil
	}

	return nil, errors.New("no information about the type of animal specified")
}

func main() {
	dogAmount, err := animal("perro")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	catAmount, err := animal("gato")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var amount float64
	amount += dogAmount(5)
	amount += catAmount(8)
	fmt.Println("Amount needed: ", amount)
}
