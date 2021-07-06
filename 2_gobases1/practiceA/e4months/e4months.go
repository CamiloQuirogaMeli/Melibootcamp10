// Realizar una aplicación que contenga una variable con el número del mes, imprimir el mes
// que corresponda. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y
// por qué?
// Ej: 7, Julio
package main

import "fmt"

var months = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func main() {
	var month int
	fmt.Printf("Ingrese el mes: ")
	fmt.Scanf("%d\n", &month)
	fmt.Println(month, ", ", months[month-1])
}
