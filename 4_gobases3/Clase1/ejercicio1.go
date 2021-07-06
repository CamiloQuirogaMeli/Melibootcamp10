package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*

		Una empresa que se encarga de vender productos de limpieza necesita:
			1. Implementar una funcionalidad para guardar un archivo de texto, con la información
			de productos comprados, separados por punto y coma.
			2. Debe tener el id del producto, precio y la cantidad.
			3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.

	*/
	var nombreArchivo string
	var contenido string
	var opcion int
	salir := false
	var id int
	var precio float64
	var cantidad int

	fmt.Println("Digita el nombre del archivo: ")
	fmt.Scanln(&nombreArchivo)

	for !salir {

		fmt.Println("Digita una opción:")
		fmt.Println("\t 1: Añadir un nuevo producto")
		fmt.Println("\t 2: Ver los productos añadidos hasta el momento")
		fmt.Println("\t 3: Guardar en un txt los productos")
		fmt.Println("\t 4: Salir")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Digita el ID:")
			fmt.Scanln(&id)
			fmt.Println("Digita el precio:")
			fmt.Scanln(&precio)
			fmt.Println("Digita la cantidad:")
			fmt.Scanln(&cantidad)

			contenido = fmt.Sprint(contenido, id, ",", precio, ",", cantidad, ";")
		case 2:
			if len(contenido) > 0 {
				fmt.Println("Los productos añadidos hasta el momento son:")
				fmt.Println("ID\tPrecio\tCantidad")
				fmt.Println(contenido)
			} else {
				fmt.Println("No hay productos aún")
			}
		case 3:
			if len(contenido) > 0 {
				guardarComoTxt(contenido, nombreArchivo)
			} else {
				fmt.Println("Debes añadir un producto primero")
			}
		case 4:
			salir = true
		}

	}

}

func guardarComoTxt(contenido string, nombre string) {
	d1 := []byte(contenido)
	nombreArchivo := "./" + nombre + ".txt"
	err := ioutil.WriteFile(nombreArchivo, d1, 0644)
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println("El arhivo se guardó correctamente")
	}
}
