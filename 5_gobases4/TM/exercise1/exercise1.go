package main

import (
	"fmt"
	"log"
	"os"
)

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error: %v", e.msg)
}

func main() {
	salary := 10000
	myError := MyError{msg: "el salario ingresado no alcanza el mínimo imponible"}
	if salary < 150000 {
		log.Fatal(myError.Error())
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto.")
}
