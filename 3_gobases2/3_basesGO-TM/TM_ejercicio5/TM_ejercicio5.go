package main

import (
	"errors"
	"fmt"
)

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

ejemplo:
const (
perro = "perro"
gato = "gato"
)

...
animalPerro, err := Animal(perro)
animalGato, err := Animal(gato)

...
var cantidad float64
cantidad += animalPerro(5)
cantidad += animalGato(8)


*/

const (
	tarantula = "tarantula"
	hamster   = "hamster"
	perro     = "perro"
	gato      = "gato"
)

/*Menu de tipo de animales */
func animal(tipo string) (func(cant int) int, error) {

	switch tipo {
	case tarantula:
		return calculaTarantula, nil
	case hamster:
		return calculaHamster, nil
	case perro:
		return calculaPerro, nil
	case gato:
		return calculaGato, nil
	default:
		return calculaTarantula, errors.New(("el tipo especificado no es válido"))
	}

}

/*Calculo la cantidad de alimento en g para tarantulas*/
func calculaTarantula(cant int) int {
	if cant < 0 {
		return 0
	} else {
		return cant * 150
	}
}

/*Calculo la cantidad de alimento en g para hamsters*/
func calculaHamster(cant int) int {
	if cant < 0 {
		return 0
	} else {
		return cant * 250
	}
}

/*Calculo la cantidad de alimento en kg para perros*/
func calculaPerro(cant int) int {
	if cant < 0 {
		return 0
	} else {
		return cant * 10
	}
}

/*Calculo la cantidad de alimento en kg para gatos*/
func calculaGato(cant int) int {
	if cant < 0 {
		return 0
	} else {
		return cant * 5
	}
}

func main() {

	/*PERRO*/
	funPerro, err := animal(perro)
	cantPerro := 10
	kgPerro := 0
	if err == nil {
		kgPerro = funPerro(cantPerro)
	}

	fmt.Println("Cantidad de perros: ", cantPerro)
	fmt.Println("Kg de alimento: ", kgPerro)

	/*GATO*/
	funGato, err := animal(gato)
	cantGato := 4
	kgGato := 0
	if err == nil {
		kgGato = funGato(cantGato)
	}

	fmt.Println("Cantidad de gatos: ", cantGato)
	fmt.Println("Kg de alimento: ", kgGato)

	/*HAMSTERS*/
	funHamster, err := animal(hamster)
	cantHamster := 2
	gHamster := 0
	if err == nil {
		gHamster = funHamster(cantHamster)
	}

	fmt.Println("Cantidad de hamsters: ", cantHamster)
	fmt.Println("Gramos de alimento: ", gHamster)

	/*TARANTULAS*/
	funTarantula, err := animal(tarantula)
	cantTarantula := 1
	gTarantula := 0
	if err == nil {
		gTarantula = funTarantula(cantTarantula)
	}

	fmt.Println("Cantidad de tarantulas: ", cantTarantula)
	fmt.Println("Gramos de alimento: ", gTarantula)

}
