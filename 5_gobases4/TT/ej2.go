package main

import (
	"errors"
	"fmt"
	"os"
)

type Cliente struct {
	Legajo    int
	NomYApe   string
	DNI       string
	Tel       string
	Domicilio string
}

func leerArchivo(path string, c1 chan string) {
	defer func() {
		err := recover()
		if err != nil {
			c1 <- "El archivo indicado no fue encontrado o esta dañado"
		}
	}()
	_, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	c1 <- "Archivo leido"
}

func (c Cliente) validarCliente(chanell chan string) {
	var err error
	if c.Legajo == 0 {
		err = errors.New("error: el numero de legajo no debe ser zero")
	}
	if c.NomYApe == "" {
		err = errors.New("error: el nombre no debe ser zero")

	}
	if c.Domicilio == "" {
		err = errors.New("error: el domicilio no debe ser zero")

	}
	if c.Tel == "" {
		err = errors.New("error: el telefono no debe ser zero")

	}
	if c.DNI == "" {
		err = errors.New("error: el dni no debe ser zero")
	}
	defer func() {
		err := recover()
		if err != nil {
			chanell <- "ha ocurrido un error con el tipo de datos seteado"
		}
	}()
	if err != nil {
		panic(err)
	}
	chanell <- "Cliente ok"
}

func main() {
	path := "./customers.txt"
	cli := Cliente{10, "Daniel Sanchez", "398429384", "3583982", "Lavalle 343"}

	chanell := make(chan string)
	c1 := make(chan string)

	go cli.validarCliente(chanell)
	go leerArchivo(path, c1)

	var resValidarCliente = "ok"
	var resLeerArchivo = "ok"
	resValidarCliente = <-chanell
	resLeerArchivo = <-c1

	// fmt.Println("Los tipos de datos son correctos?", ok)
	fmt.Println(resValidarCliente)
	fmt.Println(resLeerArchivo)

	fmt.Println("Fin de la ejecucion")
}

//resolver con go rutines
