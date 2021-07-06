package main

import (
	"fmt"
)

type MyError struct {
	msj string
}

func (err *MyError) Error() string {
	return err.msj
}

func main() {

	myError := MyError{"error: el salario ingresado no alcanza el mínimo imponible"}
	salary := 15000

	if salary < 150000 {
		fmt.Println(myError.Error())
	} else {
		fmt.Println("\"Debe pagar impuesto\"")
	}

}
