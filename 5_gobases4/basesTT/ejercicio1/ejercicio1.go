package main

import (
	"fmt"
	"io/ioutil"
)

func leerArchivo() {
	defer func() {
		errs := recover()
		if errs != nil {
			fmt.Println("error: el archivo indicado no fue encontrado o esta dañado")
		}
	}()
	_, err := ioutil.ReadFile("./customers.txt")
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Iniciando lectura")
	leerArchivo()
	fmt.Print("ejecucion finalizada")
}
