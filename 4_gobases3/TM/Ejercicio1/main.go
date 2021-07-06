package main

import (
	"fmt"
	"os"
)

func main() {
	ejercicio1()
}

func ejercicio1() {
	textoEjercicio1 := []byte("123;1000.13;1\n124;2541.31;2\n125;348.25;1")
	err := os.WriteFile("./ejercicio1", textoEjercicio1, 0644)

	if err == nil {
		fmt.Println("Se genero el archivo correctamente")
	} else {
		fmt.Println(err)
	}
}
