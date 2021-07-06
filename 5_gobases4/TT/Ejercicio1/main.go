package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Ejecucion Terminada!")

	leerArchivo()
}

func leerArchivo() string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	abirArchivo, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("No se encontro ningún archivo disponible")
	}
	return string(abirArchivo)
}
