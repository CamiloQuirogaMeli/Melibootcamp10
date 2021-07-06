package main
import (
	"fmt"
	"errors"
)

/* Ejercicio 5
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
cantidad += animalGato(8) */

const (
	perro = "perro"
	gato = "gato"
	hamster = "hamster"
	tarantula = "tarantula"
)

func main() {
	perro, err := Animal(perro)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(perro(3)/1000, "KG")
	}

	gato, err := Animal(gato)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(gato(3)/1000, "KG")
	}

	hamster, err := Animal(hamster)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(hamster(3)/1000, "KG")
	}

	tarantula, err := Animal(tarantula)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(tarantula(3)/1000, "KG")
	}
}

func Animal(a string) (func (cantidad int) float64, error) {
	switch a {
		case perro: {
			return easy(10000.0), nil
		}
		case gato: {
			return easy(5000.0), nil
		}
		case hamster: {
			return easy(250.0), nil
		}
		case tarantula: {
			return easy(150.0), nil
		}
		default: {
			return nil, errors.New("Error. Mascota desconocida.")
		}
	}
}

func easy(gramos float64) (func (cantidad int) float64) {
	return func (cantidad int) float64 {
		return float64(cantidad) * gramos
	}
}