package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	defer func() {
		fmt.Println("Ejecución finalizada")
	}()

	content, err := ioutil.ReadFile("./customers.txt")

	if err != nil {
		panic("El archivo indicado no fue encontrado o esta dañado")
	} else {
		fmt.Println(content)
	}
}
