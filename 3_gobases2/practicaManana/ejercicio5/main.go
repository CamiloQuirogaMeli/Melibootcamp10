package main

import (
	"errors"
	"fmt"
)

/*Un refugio necesita calcular cuanto alimento debe comprar para las mascotas, por el
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
cantidad del tipo de animal especificado.*/

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

var comidaAnimal float64

func main() {
	var animal string
	var cant int
	fmt.Println("Ingrese animal: ")
	fmt.Scanln(&animal)
	fmt.Println("Ingrese cantidad de ", animal, " que tiene el refugio:")
	fmt.Scanln(&cant)

	cantAlimento, err := buscarAnimal(animal, cant)
	if cantAlimento == -1 {
		fmt.Println(err)
	} else {
		fmt.Println("Se debe combrar ", cantAlimento, " kilos")
	}

}

func buscarAnimal(animal string, cantidadAnimal int) (float64, error) {

	switch animal {
	case perro:
		comidaAnimal = 10
		return calcularCantidad(cantidadAnimal), nil
	case gato:
		comidaAnimal = 5
		return calcularCantidad(cantidadAnimal), nil
	case hamster:
		comidaAnimal = 0.25
		return calcularCantidad(cantidadAnimal), nil
	case tarantula:
		comidaAnimal = 0.15
		return calcularCantidad(cantidadAnimal), nil
	}
	return -1, errors.New("No existe animal")
}

func calcularCantidad(cant int) float64 {
	return float64(cant) * comidaAnimal
}
