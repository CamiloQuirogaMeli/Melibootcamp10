package main

import (
	"errors"
	"fmt"
)

func main() {
	var option int
	var aux int = 1
	for aux != 0 {

		const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

		const options string = "1. Iniciar \n 0. Salir"

		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			var animal string
			fmt.Println("Ingresa el nombre del animal")
			fmt.Scan(&animal)
			functionAction, err := evaluateAnimal(animal)

			if err != nil {
				fmt.Println(err)
				break
			}

			var nAnimal int
			fmt.Println("Ingresa el numero de animales")
			fmt.Scan(&nAnimal)

			response := functionAction(nAnimal)
			fmt.Println("*****RESPUESTA***\n", response)

		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}
	}

}

func evaluateAnimal(animal string) (func(value int) string, error) {

	const (
		DOG     = "perro"
		CAT     = "gato"
		HAMSTER = "hamster"
		SPIDER  = "tarantula"
	)

	switch animal {
	case DOG:
		return getDogFood, nil
	case HAMSTER:
		return getHamsterFood, nil
	case SPIDER:
		return getSpiderFood, nil
	case CAT:
		return getCatFood, nil
	default:
		return nil, errors.New("¡¡ERROR!! has ingresado un animal invalido")
	}
}

func getDogFood(nAnimal int) string {
	var amountFood int = 10
	return "Se necesitan " + fmt.Sprint(nAnimal*amountFood) + "kg de alimento"
}
func getCatFood(nAnimal int) string {
	var amountFood int = 5
	return "Se necesitan " + fmt.Sprint(nAnimal*amountFood) + "kg de alimento"
}
func getSpiderFood(nAnimal int) string {
	var amountFood int = 150
	return "Se necesitan " + fmt.Sprint(nAnimal*amountFood) + "gramos de alimento"
}
func getHamsterFood(nAnimal int) string {
	var amountFood int = 250
	return "Se necesitan " + fmt.Sprint(nAnimal*amountFood) + "gramos de alimento"
}
