package main

import (
	e "errors"
	f "fmt"
)

const (
	DOG       = "dog"
	CAT       = "cat"
	HAMSTER   = "hamster"
	TARANTULA = "tarantula"
)

func calcFoodDog(dog float32) float32 {
	if dog > 0 {
		return dog * 10
	} else {
		return dog
	}
}

func calcFoodCat(cat float32) float32 {
	if cat > 0 {
		return cat * 5
	} else {
		return cat
	}
}

func calcFoodHamster(hamster float32) float32 {
	if hamster > 0 {
		return hamster * 0.25
	} else {
		return hamster
	}
}

func calcFoodTarantula(tarantula float32) float32 {
	if tarantula > 0 {
		return tarantula * 0.15
	} else {
		return tarantula
	}
}

func animal(animal string) (func(animal float32) float32, error) {
	switch animal {
	case DOG:
		return calcFoodDog, nil
	case CAT:
		return calcFoodCat, nil
	case HAMSTER:
		return calcFoodHamster, nil
	case TARANTULA:
		return calcFoodTarantula, nil
	default:
		msg := animal + " does not exist"
		return nil, e.New(msg)
	}
}

func main() {
	var total float32 = 0
	var _errors []error

	foodDog, err1 := animal(DOG)
	catFood, err2 := animal(CAT)
	hamsterFood, err3 := animal(HAMSTER)
	tarantulaFood, err4 := animal(TARANTULA)
	
	switch {
	case err1 != nil:
		_errors = append(_errors, err1)
	case err2 != nil:
		_errors = append(_errors, err2)
	case err3 != nil:
		_errors = append(_errors, err3)
	case err4 != nil:
		_errors = append(_errors, err4)
	default:
		break
	}

	if len(_errors) > 0 {
		for i, err := range _errors {
			f.Println("ERROR #:", i+1, err)
		}
	} else {
		total += foodDog(3)
		total += catFood(5)
		total += hamsterFood(8)
		total += tarantulaFood(8)

		f.Println("You need buy", total, "Kg of food for all pets")
	}
}
