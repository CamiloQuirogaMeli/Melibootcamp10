package main

import "fmt"

func main() {
	// Formas de resolverlos:
	// Utilizando un array -> Debería acceder a las posiciones restando uno
	// Utilizando un archivo con formato k/v (numMes;mes)
	// Utilizando un switch
	// Utilizando ifs/else
	// Utilizando un map
	// Elección:
	// Elegí el map porque me parece la opción que mejor se adecúa al problema
	// dada la naturaleza del mismo. Resulta simple verlo como una estructura de k/v
	// y también sencillo de leer y acceder a los elementos sumado a la claridad
	// del código.
	// Mi segunda opción sería el switch porque resulta simple y claro de ver
	// la intención del código, es decir, para cada valor, corresponde un mes a mostrar.
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	var monthNum int
	fmt.Print("Enter a month number: ")
	fmt.Scanf("%d", &monthNum)
	fmt.Println("The month is: ", months[monthNum])
}
