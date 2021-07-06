package main

import "fmt"

//Realizar una aplicación que contenga una variable con el número del mes, imprimir el mes que corresponda.
//¿Se te ocurre si se puede resolver de más de una manera?
//--> En mi opnión habrían dos formas...una es haciendo uso de un map, con el valor del mes e imprimir buscando el elemento. La otra opción, es con un array, donde el número
//del mes corresponde a la posición en el arreglo
//¿Cuál elegirías y por qué?
//--> Uso un map por facilidad

func main()  {
	var months = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	fmt.Println(months)
	fmt.Println("Ejemplo: ")
	fmt.Println("7, ", months[7])
}