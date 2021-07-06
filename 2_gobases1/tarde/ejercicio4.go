package main

import "fmt"

func main() {
	//calendario := map
	var monthsMap = map[int]string{1: "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre"}

	var monthNumber = 7

	fmt.Println("El mes número", monthNumber, "es el mes:", monthsMap[monthNumber])

}

// Otra forma de resolver este ejercicio seria usando condicionales para validar el número y a partir de este imprimir el mes correspondiente
// Otra forma sería usando una estructura de control Switch
// También hubieramos podido guardar los meses en un array de strings de forma ordenada y encontrar el nombre del mes usando el número de mes como la posición (mes = lista_meses[numero_mes - 1])
// La forma como se implementó finalmente fue usando un mapa, en donde los números son las llaves y el nombre de los meses los valores.
// Creo que la mejor forma fue con los mapas, en mi opinión es la que queda más clara (el código es más facil de entender) y los mapas son eficientes.
