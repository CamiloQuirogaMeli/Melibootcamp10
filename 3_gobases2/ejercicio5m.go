package main

import "errors"

// Un refugio necesita calcular cuanto alimento debe comprar para las mascotas, por el
// momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan haber
// muchos más animales que refugiar.
// Por perro necesitan 10 kg de alimento
// Por gato 5 kg
// Por Hamsters 250 gramos.
// Por Tarántula 150 gramos.
// Para ellos necesitamos:
// - Implementar una función Animal que reciba como parámetro un valor de tipo texto
// con el animal especificado y nos retorne una función y un error (en caso que no exista
// el animal)
// - Una función para cada animal que calcule la cantidad de alimento en base a la
// cantidad del tipo de animal especificado.

func calcFoodPet(animal string) (func() string, error) {
	switch animal {
	case "perro":
		return perro, nil
	case "gato":
		return gato, nil
	case "hamster":
		return hamster, nil
	case "tarantula":
		return tarantula, nil
	default:
		return nil, errors.New("No tenemos comida para ese animal 0.0")
	}

}

func perro() string {
	return "Para un perro necesitas 10000grs de Alimento"
}

func gato() string {
	return "Para un perro necesitas 5000grs de Alimento"

}

func hamster() string {
	return "Para un perro necesitas 250grs de Alimento"

}

func tarantula() string {
	return "Para un perro necesitas 150grs de Alimento"
}
