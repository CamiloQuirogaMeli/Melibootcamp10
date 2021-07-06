package main

import (
	"fmt"
	"os"
)

func leerArchivo(path string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	datosBytes, err := os.Open(path)
	if err != nil {
		panic("El archivo no existe o esta roto")
	}
	fmt.Println(datosBytes)
}
func main() {
	path := "./customers.txt"
	leerArchivo(path)
	fmt.Println("Fin del programa")
}
