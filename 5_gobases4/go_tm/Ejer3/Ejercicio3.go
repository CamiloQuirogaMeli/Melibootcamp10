/*
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de
error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible
(el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y
el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var salary int

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite su salario")
	scanner.Scan()

	salary, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Bad input")
		os.Exit(0)
	}
	if salary < 15000 {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
		fmt.Println(err)
	} else {
		err := fmt.Errorf("Debe pagar impuesto porque su salario es de %v", salary)
		fmt.Println(err)
	}

}
