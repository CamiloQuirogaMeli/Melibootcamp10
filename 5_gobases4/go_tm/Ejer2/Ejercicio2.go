/*
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en
reemplazo de “Error()”, se implemente “errors.New()”.
*/

package main

import (
	"bufio"
	"errors"
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
		err := errors.New("error: el salario ingresado no alcanza el mínimo imponible")
		fmt.Println(err)

	} else {
		err := errors.New("Debe pagar impuesto")
		fmt.Println(err)
	}

}
