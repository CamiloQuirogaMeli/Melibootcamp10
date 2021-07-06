package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	nombreArchivo := ""
	fmt.Println("Ingrese el nombre del archivo")
	fmt.Scanln(&nombreArchivo)
	defer func() {
		fmt.Println("-----Ejecucion finalizada-------")
	}()
	contenidoArchivo, err := ioutil.ReadFile("./" + nombreArchivo)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Println(contenidoArchivo)
	}
}
