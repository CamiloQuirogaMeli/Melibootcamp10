package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	defer func() {
		fmt.Println("Ejecución finalizada")
	}()

	content, err := ioutil.ReadFile("../Ejercicio1/Pruebita")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Print(content)
	}
}
