/* Ejercicio 2 - Impuestos de salario #2
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en
reemplazo de “Error()”, se implemente “errors.New()”. */

package main

import (
	"fmt"
	"errors"
)

func main() {
	salary := 135000

	if salary < 150000 {
		err := errors.New("error: el salario ingresado no alcanza el mínimo imponible")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}