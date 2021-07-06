package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	Read()
	fmt.Println("ejecucion finaliza!")
}

func Read() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := ioutil.ReadFile("productos.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}
}
