package main

import (
	"fmt"
)

func ej2() {
	var temperatura, humedad, presionAtm int = 18, 32, 1012
	const UBICACION = "Guaymallén, Mendoza, Argentina"
	fmt.Println("La temperatura en: "+UBICACION+" es de: ", temperatura)
	fmt.Println("La humedad es de: ", humedad) // se debe poner la coma para ingresar en un string un int
	fmt.Println("La presión atmosférica es de: ", presionAtm)
}
