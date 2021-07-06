package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// leer el arreglo de bytes del archivo
	ejercicio1, err := ioutil.ReadFile("../Ejercicio1/ejercicio1.txt")
	if err != nil {
		log.Fatal(err)
	}
	datosComoString := string(ejercicio1)
	var datosComoString2 string
	for _, valor := range datosComoString {
		if valor == 59 {
			datosComoString2 = datosComoString2 + "\t\t"
		} else {

			datosComoString2 = datosComoString2 + string(valor)
		}
	}

	// convertir el arreglo a string
	//datosComoString := string(ejercicio1)
	// imprimir el string
	fmt.Println(datosComoString2)
}
