package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func main() {
	defer func() {
		fmt.Println("Ejecucion finalizada")
	}()

	filename := "../customers.txt"
	file, err := readFile(filename)

	if err != nil {
		panic("El archivo indicado no fue encontrado o esta dañado")
	}
	fmt.Println(string(file))
}
