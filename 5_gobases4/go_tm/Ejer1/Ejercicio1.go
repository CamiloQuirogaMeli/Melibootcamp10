/*
1. En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo
“int”.
2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje
“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe
pagar impuesto”.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type error interface {
	Error() string
}

type CustomError struct {
	status int
	msg    string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

func CustomErrorTest(status int) (int, error) {
	if status >= 300 {
		return 400, &CustomError{
			status: status,
			msg:    "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return 200, nil
}

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
		_, err := CustomErrorTest(500)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		CustomErrorTest(200)
		fmt.Println("Debe pagar impuesto")
	}

}
