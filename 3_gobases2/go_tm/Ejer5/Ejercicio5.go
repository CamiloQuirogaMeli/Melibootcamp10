/*
Un refugio necesita calcular cuanto alimento debe comprar para las mascotas, por el
momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan haber
muchos más animales que refugiar.

	Por perro necesitan 10 kg de alimento
	Por gato 5 kg
	Por Hamsters 250 gramos.
	Por Tarántula 150 gramos.
Para ellos necesitamos:
	- Implementar una función Animal que reciba como parámetro un valor de tipo texto
	con el animal especificado y nos retorne una función y un error (en caso que no exista
	el animal)
	- Una función para cada animal que calcule la cantidad de alimento en base a la
	cantidad del tipo de animal especificado.
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	PERRO     = "perros"
	GATO      = "gatos"
	HAMSTER   = "hamsters"
	TARANTULA = "tarantulas"
	OPTIONS   = `
	1. Si desea calcular la comida para perros digite "perros"
	2. Si desea calcular la comida para gatos digite "gatos"
	3. Si desea calcular la comida para hamsters digite "hamsters"
	4. Si desea calcular la comida para tarantula digite "tarantulas"
	5. Si desea salir digite "salir"
	`
)

func main() {
	continuar := true

	for continuar {
		println("Por favor siga las siguientes recomendaciones")
		println(OPTIONS)
		print("->")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		typeAnimal := animal(scanner.Text())
		println("¿Cuantos ", scanner.Text(), "son?")
		var quantity int
		_, errQuantity := fmt.Scanf("%d", &quantity)

		if errQuantity != nil {
			fmt.Println("Bad Input")
			os.Exit(0)
		}
		result, errResult := typeAnimal(quantity)

		if errResult != nil {

			fmt.Println(errResult)
		} else {
			var unity string
			if scanner.Text() == "perros" || scanner.Text() == "gatos" {
				unity = "KG"
			} else {
				unity = "G"
			}

			println("La cantidad de comida es: ", strconv.FormatFloat(result, 'f', -2, 64), unity)
		}

		println()
	}
}

func animal(opcion string) func(quantity int) (float64, error) {
	switch opcion {
	case PERRO:
		return opPerro
	case GATO:
		return opGato
	case HAMSTER:
		return opHamster
	case TARANTULA:
		return opTarantula
	case "salir":
		println("Gracias por usar el programa")
		os.Exit(0)
	default:
		println("Opcion invalida")
		os.Exit(0)
	}

	return nil
}

func opPerro(q int) (float64, error) {

	if math.Signbit(float64(q)) {

		return 0, errors.New("No se aceptan valores negativos")
	}
	return float64(q) * 10, nil
}
func opGato(q int) (float64, error) {
	if math.Signbit(float64(q)) {
		return 0, errors.New("No se aceptan valores negativos")
	}
	return float64(q) * 5, nil
}
func opHamster(q int) (float64, error) {
	if math.Signbit(float64(q)) {
		return 0, errors.New("No se aceptan valores negativos")
	}
	return float64(q) * 250, nil
}
func opTarantula(q int) (float64, error) {
	if math.Signbit(float64(q)) {
		return 0, errors.New("No se aceptan valores negativos")
	}
	return float64(q) * 150, nil
}
