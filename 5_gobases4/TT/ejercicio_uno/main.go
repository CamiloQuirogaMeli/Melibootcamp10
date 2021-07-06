package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	defer func() {

		err := recover()

		if err != nil {
			fmt.Println(err)
		}

	}()

	_, err := ioutil.ReadFile("customers.txt")

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}

func main() {

	readFile()

	fmt.Println("ejecucion finalizada")
}
