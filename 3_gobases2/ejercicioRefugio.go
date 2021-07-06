package main

import "fmt"
import "errors"

//Un refugio necesita calcular cuanto alimento debe comprar para las mascotas, por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan haber
//muchos más animales que refugiar.

//Por perro necesitan 10 kg de alimento
//Por gato 5 kg
//Por Hamsters 250 gramos.
//Por Tarántula 150 gramos.

//Para ellos necesitamos:
//- Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y nos retorne una función y un error (en caso que no exista
//el animal)
//- Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

const (
	perro = "perro"
	gato = "gato"
	hamster = "hamster"
	tarantula = "tarantula"
)

func main()  {
	var cantidad float64

	animalPer, err1 := Animal(perro)
	animalGat, err2 := Animal(gato)
	animalHam, err3 := Animal(hamster)
	animalTar, err4 := Animal(tarantula)
	_, err5 := Animal("na")

	if(err1 != nil){
		fmt.Println("ERROR:", err1)
	} else {
		cantidad += animalPer(5)
	}

	if(err2 != nil){
		fmt.Println("ERROR:", err2)
	} else {
		cantidad += animalGat(8)
	}

	if(err3 != nil){
		fmt.Println("ERROR:", err3)
	} else {
		cantidad += animalHam(3)
	}

	if(err4 != nil){
		fmt.Println("ERROR:", err4)
	} else {
		cantidad += animalTar(1)
	}

	if(err5 != nil){
		fmt.Println("ERROR:", err5)
	} else {
		fmt.Println(nil)
	}
	
	fmt.Println("La cantidad total de alimento que debe comprar el refugio es", cantidad, "Kg")
}

func Animal(an string) (func(can float64) float64, error) {
	switch an {
	case perro:
		return canPerr, nil
	case gato:
		return canGat, nil
	case hamster:
		return canHams, nil
	case tarantula:
		return canTar, nil
	}

	return nil, errors.New("El animal que desea calcular no existe.")
}

func canPerr(can float64) float64 {
	return can * 10
}

func canHams(can float64) float64 {
	return can * (250 / 1000)
}

func canGat(can float64) float64 {
	return can * 5
}

func canTar(can float64) float64 {
	return can * (150 / 1000)
}
