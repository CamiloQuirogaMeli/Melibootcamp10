package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	/*
			La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima
		por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID
		y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe
		visualizar el total (Sumando precio por cantidad)
	*/
	var nombreArchivo string
	salir := false
	var opcion int

	for !salir {
		fmt.Println("Digita una opción:")
		fmt.Println("\t 1: Leer un archivo")
		fmt.Println("\t 2: Salir")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Digita el nombre del archivo para leer:")
			fmt.Scanln(&nombreArchivo)
			nombreArchivo = "./" + nombreArchivo + ".txt"
			dat, err := ioutil.ReadFile(nombreArchivo)
			if err != nil {
				fmt.Println("Nombre no encontrado")
			} else {

				ImprimirFormato(string(dat))
			}
		case 2:
			salir = true
		}
	}
}

func ImprimirFormato(datos string) {
	lineas := strings.Split(datos, ",")
	var espacios int

	fmt.Println(lineas)
	fmt.Println("ID       Precio   cantidad ")

	for i := 0; i < len(lineas); i += 3 {
		for j := 0; j < 3; j++ {
			espacios = 9 - len(lineas[i+j])
			fmt.Print(lineas[i+j])
			for e := 0; e < espacios; e++ {
				fmt.Print(" ")
			}
			fmt.Print("")
		}
		fmt.Println("")

	}
	fmt.Println("\t", miSuma(lineas))

}

func miSuma(numeros []string) float64 {
	var sum float64
	for _, i := range numeros {
		f, err := strconv.ParseFloat(i, 32)
		if err != nil {
			fmt.Println("Dato desconocido")
		} else {
			sum += f
		}
	}
	return sum
}
