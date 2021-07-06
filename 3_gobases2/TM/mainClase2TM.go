package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ejercicio 1 = 1 \nEjercicio 2 = 2 \nEjercicio 3 = 3\nEjercicio 4 = 4\nEjercicio 5 = 5")
	fmt.Print("Ingrese aplicación a la que desea ingresar: (1 o 2 o 3 o 4 o 5)")
	var ejercicio int
	fmt.Scanln(&ejercicio)

	switch ejercicio {
	case 1:
		imp := impuesto()

		fmt.Println("El impuesto a su sueldo es de: ", imp)

	case 2:
		prom, err := promedio(9, 9, -10, 2)
		if err != nil {
			fmt.Println(prom, err)
		} else {
			fmt.Println(prom, err)
		}

	case 3:
		salario()
	case 4:
		fmt.Print("Ingrese funcion a la que desea ingresar: (minimo, promedio, maximo)")
		var funcionAplicar string
		fmt.Scanln(&funcionAplicar)
		//var valores = []float64{2, 3, 3, 4, 1, 2, 4, 5}
		//operacion(funcionAplicar, 2, 3, 3, 4, 1, 2, 4, 5)

		_, err := operacion(funcionAplicar)

		if funcionAplicar == MINIMO {
			minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		} else if funcionAplicar == PROMEDIO {
			promFunc(2, 3, 3, 4, 1, 2, 4, 5)
		} else if funcionAplicar == MAXIMO {
			maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		} else {
			fmt.Println(err)
		}

	case 5:
		fmt.Print("Ingrese animal a comprar comida: (perro, gato, tarantula, hamster)\n")
		var animalIngresado string
		fmt.Scanln(&animalIngresado)
		fmt.Print("Ingrese cantidad de animales: \n")
		var cantidad float64
		fmt.Scanln(&cantidad)

		_, err := comidaPorAnimal(animalIngresado)

		if animalIngresado == perro {
			res := comidaPerro(cantidad)
			fmt.Println("La cantidad de comida para el animal: ", animalIngresado, "es :", res, " kilos.")
		} else if animalIngresado == gato {
			res := comidaGato(cantidad)
			fmt.Println("La cantidad de comida para el animal: ", animalIngresado, "es :", res, " kilos.")
		} else if animalIngresado == hamster {
			res := comidaHamster(cantidad)
			fmt.Println("La cantidad de comida para el animal: ", animalIngresado, "es :", res, " gramos.")
		} else if animalIngresado == tarantula {
			res := comidaTarantula(cantidad)
			fmt.Println("La cantidad de comida para el animal: ", animalIngresado, "es :", res, " gramos. ")
		} else {
			fmt.Println(err)
		}

	case 6:
		break

	}

}
