package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
	pezGlobo  = "pezGlobo"
)

func printError(err ...error) {
	for _, err := range err {
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}
func animalDog(quantity int) float64 {
	return float64(quantity) * 10.0
}
func animalCat(quantity int) float64 {
	return float64(quantity) * 5.0
}
func animalHamster(quantity int) float64 {
	return float64(quantity) * 0.25
}
func animalTarantula(quantity int) float64 {
	return float64(quantity) * 0.15
}
func Animal(name string) (func(int) float64, error) {
	switch name {
	case "perro":
		{
			return animalDog, nil
		}
	case "gato":
		{
			return animalCat, nil
		}
	case "hamster":
		{
			return animalHamster, nil
		}
	case "tarantula":
		{
			return animalTarantula, nil
		}
	default:
		return nil, errors.New("no existe el animal solicitado : " + name)
	}
}

func main() {
	animalPerro, err1 := Animal(perro)
	animalGato, err2 := Animal(gato)
	animalHamster, err3 := Animal(hamster)
	animalTarantula, err4 := Animal(tarantula)
	//animalPezGlobo, err5 := Animal(pezGlobo)
	var cantidad float64
	printError(err1, err2, err3, err4)
	cantidad += animalPerro(5)
	cantidad += animalGato(8)
	cantidad += animalHamster(10)
	cantidad += animalTarantula(3)
	//cantidad += animalPezGlobo(10)
	fmt.Printf("La cantidad de alimento es: %.2f kg", cantidad)
}
