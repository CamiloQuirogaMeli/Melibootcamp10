package main

import "fmt"

func main() {
	var numero int
	var meses [12]string = [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	fmt.Println("Numero de mes: ")
	fmt.Scanln(&numero)
	fmt.Println(meses[numero-1])
	// Existen diversas formas de realizar éste ejercicio. Utilizando slice y map.
	// Como la lista de meses no varía, decidí realizarlo con array por una cuestión de
	// simplicidad, performance, y porque no son utiles las capacidades que dan las demás estructuras
}
