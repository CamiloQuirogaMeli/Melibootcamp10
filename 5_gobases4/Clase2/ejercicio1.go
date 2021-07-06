package main

import (
	"fmt"
	"os"
)

func leerArchivo(nombreArchivo string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("el archivo indicado no fue encontrado o está dañado")
			fmt.Println("Ejecución finalizada")
		}
	}()

	fmt.Println("Se intentará leer el archivo")

	dat, err := os.Open(nombreArchivo)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(dat)
	}
}

func main() {

	/*

		Un estudio contable necesita acceder a los datos de sus empleados para poder realizar
		distintas liquidaciones. Para ello, cuentan con todo el detalle necesario en un archivo .txt.
		1. Es necesario desarrollar la funcionalidad para poder leer el archivo .txt que nos indica
		el cliente, sin embargo, no han pasado el archivo a leer por nuestro programa.
		2. Desarrolla el código necesario para leer los datos del archivo llamado “customers.txt”
		(recuerda lo visto sobre el pkg “ioutil”).
		Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso,
		el programa deberá arrojar un panic al intentar leer un archivo que no existe,
		mostrando el mensaje “el archivo indicado no fue encontrado o está dañado”.
		Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.

	*/

	nombreArchivo := "customers.txt"
	salir := false
	var opcion int

	for !salir {
		fmt.Println("Digita una opción:")
		fmt.Println("\t 1: Leer el archivo customers.txt")
		fmt.Println("\t 2: Salir")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			leerArchivo(nombreArchivo)
		case 2:
			salir = true
		}
		fmt.Println("Ejecución finalizada")

	}

}
