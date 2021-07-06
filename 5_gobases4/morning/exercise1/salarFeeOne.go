package main

import "fmt"

const MIN_SALARY int = 150000

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error: %s", e.Msg)
}

func main() {

	var salary int = 200

	if salary < MIN_SALARY {
		err := MyError{"el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}
