package main

import (
	"errors"
	"fmt"
)

func Animal(anim string) (func (float64) float64, error){

	switch anim{
	case "perro":
		cFinal := cPerro
		return cFinal, nil
	case "gato":
		cFinal := cGato
		return cFinal, nil
	case "hamster":
		cFinal := cHamster
		return cFinal, nil
	case "tarantula":
		cFinal := cTarantula
		return cFinal, nil
	}

	return nil, errors.New("error: el animal no existe.")
}

func cPerro(c float64) float64{
	return 10*c
}

func cGato(c float64) float64{
	return 5*c
}

func cHamster(c float64) float64{
	return 0.250*c
}

func cTarantula(c float64) float64{
	return 0.150*c
}

const(
	PERRO = "perro"
	GATO = "gato"
	HAMSTER = "hamster"
	TARANTULA = "tarantula"
)

func main(){
	var cantidad float64 = 0
	animalPerro, err := Animal(PERRO)
	animalGato, err := Animal(GATO)
	animalHamster, err := Animal(HAMSTER)
	animalTarantula, err := Animal(TARANTULA)

	cantidad += animalPerro(4)
	cantidad += animalHamster(5)
	cantidad += animalGato(2)
	cantidad += animalTarantula(10)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("La cantidad final de alimento es: ", cantidad, "Kg")
	}

}

/*Un refugio necesita calcular cuanto alimento debe comprar para las mascotas, por el
momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan haber
muchos más animales que refugiar.

Por perro necesitan 10 kg de alimento
Por gato 5 kg
Por Hamsters 250 gramos.
Por Tarántula 150 gramos.

3

Para ellos necesitamos:
- Implementar una función Animal que reciba como parámetro un valor de tipo texto
con el animal especificado y nos retorne una función y un error (en caso que no exista
el animal)
- Una función para cada animal que calcule la cantidad de alimento en base a la
cantidad del tipo de animal especificado.
*/
