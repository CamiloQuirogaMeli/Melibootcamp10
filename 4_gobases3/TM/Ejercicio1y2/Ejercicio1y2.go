package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	/*
		Ejercicio 1 - Guardar archivo

		Una empresa que se encarga de vender productos de limpieza necesita:
		1. Implementar una funcionalidad para guardar un archivo de texto, con la información
		de productos comprados, separados por punto y coma.
		2. Debe tener el id del producto, precio y la cantidad.
		3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.
	*/
	/*
		texto := []byte("001 , 250, 6")
		err := ioutil.WriteFile("productos.txt", texto, 0644)

		if err != nil {
			fmt.Println("Error", err)
		}
	*/
	/*
		Ejercicio 2 - Leer archivo
		La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima
		por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID
		y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe
		visualizar el total (Sumando precio por cantidad)
	*/

	data, err := ioutil.ReadFile("productos.txt")

	if err != nil {
		fmt.Println("Error", err)
	} else {

		fmt.Println("ID \t \t \t Precio\tCantidad")

		lineas := strings.Split(string(data), ";\n")

		for i, value := range lineas {

			if i < len(value)-1 {
				exit := strings.Split(value, " ")
				fmt.Println(exit[0], "\t\t\t", exit[1], "\t", exit[2])
			}

		}
	}
}
