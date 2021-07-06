/* Ejercicio 1 - Impuestos de salario #1
1. En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo
“int”.
2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje
“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe
pagar impuesto”. */
package main

import (
	"fmt"
)

type myError struct {
	text string
}
func (e *myError) Error() string {
	return fmt.Sprintf("%s", e.text)
}

func main() {
	salary := 135000

	if salary < 150000 {
		err := &myError{"error: el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}