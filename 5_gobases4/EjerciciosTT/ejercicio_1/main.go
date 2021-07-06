package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("ejecución inicializada")

	defer func() {
		fmt.Println("ejecución finalizada")
	}()

	_, err := ioutil.ReadFile("./customers.txt")

	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}

}
