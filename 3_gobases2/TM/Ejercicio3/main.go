package main

import (
	"errors"
	"fmt"
)

func main() {
	var minutos int
	var categoria string
	var seguir bool = true
	var minutosValido bool = false
	var categoriaValida bool = false

	for seguir {
		for !minutosValido {
			fmt.Println("Ingresá los minutos trabajados: ")
			fmt.Scanln(&minutos)
			err1 := validarNumero(minutos)
			if err1 != nil {
				fmt.Println(err1)
			} else {
				minutosValido = true
			}
		}
		minutosValido = false

		for !categoriaValida {
			fmt.Println("Ingresá la categoría (A, B o C): ")
			fmt.Scanln(&categoria)
			err := validarCategoria(categoria)
			if err != nil {
				fmt.Println(err)
			} else {
				categoriaValida = true
			}
		}
		categoriaValida = false

		salario := calcularSalario(minutos, categoria)
		fmt.Println("El salario es: ", salario)

		fmt.Println("¿Querés probar otro salario? Ingresa true o false")
		fmt.Scanln(&seguir)

	}

}

func calcularSalario(minutos int, categoria string) (salario float64) {
	horas := calcularHoras(minutos)

	switch categoria {
	case "A":
		salario = float64(3000) * float64(horas)
		salario = salario * 1.5
	case "B":
		salario = float64(1500) * float64(horas)
		salario = salario * 1.2
	case "C":
		salario = float64(1000) * float64(horas)
	}
	return
}

func calcularHoras(minutos int) (horas int) {
	horas = minutos / 60
	return
}

func validarNumero(numero int) (err error) {
	if numero < 0 {
		err = errors.New("el número debe ser mayor a 0")
	}
	return
}

func validarCategoria(letra string) (err error) {
	if letra == "A" || letra == "B" {
	} else if letra == "C" {
	} else {
		err = errors.New("la categoría debe ser A, B o C")
	}
	return
}
