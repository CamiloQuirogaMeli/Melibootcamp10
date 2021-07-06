// Ejercicio 4 - A qué mes corresponde

// Realizar una aplicación que contenga una variable con el número del mes, imprimir el mes
// que corresponda. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y
// por qué?
// Ej: 7, Julio

package main

import (
	"fmt"
	"time"
)

func main() {

	var month int

	fmt.Printf("Ingrese el número de mes: ")
	fmt.Scan(&month)

	for month > 12 || month < 1 {
		fmt.Printf("Error en el ingreso del mes, reingrese")
		fmt.Scan(&month)
	}

	switch month {
	case 1:
		fmt.Printf("El mes que indico es Enero")
	case 2:
		fmt.Printf("El mes que indico es Febrero")
	case 3:
		fmt.Printf("El mes que indico es Marzo")
	case 4:
		fmt.Printf("El mes que indico es Abril")
	case 5:
		fmt.Printf("El mes que indico es Mayo")
	case 6:
		fmt.Printf("El mes que indico es Junio")
	case 7:
		fmt.Printf("El mes que indico es Julio")
	case 8:
		fmt.Printf("El mes que indico es Agosto")
	case 9:
		fmt.Printf("El mes que indico es Septiembre")
	case 10:
		fmt.Printf("El mes que indico es Octumbre")
	case 11:
		fmt.Printf("El mes que indico es Noviembre")
	case 12:
		fmt.Printf("El mes que indico es Diciembre")
	}
	fmt.Println()

	// Otra forma de hacerlo, si es que un usuario no debe de ingresar el valor del mes,
	// es tomarlo de la variable de sistema e imprimirla

	fmt.Println("El mes en el que nos encontramos es:", time.Now().Local().Month())

	// Otra forma puede ser con los maps para el ingreso del valor por medio del usuario

	fmt.Printf("Indicar el mes con mapas: ")

	var months = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto",
		9: "Septiembre", 10: "Octumbre", 11: "Noviembre", 12: "Diciembre"}

	fmt.Scan(&month)
	for month > 12 || month < 1 {
		fmt.Printf("Error en el ingreso del mes, reingrese")
		fmt.Scan(&month)
	}

	fmt.Printf("El mes que indico es: %s \n", months[month])

	//Creo que la más sencilla es con un mapa ya que seteamos directamente los meses y segun el número solo utilizamos el elemento de la key
}
