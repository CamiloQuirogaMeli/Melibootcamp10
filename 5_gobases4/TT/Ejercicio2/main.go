package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	c := cliente{
		Legajo:         generarNumeroLegajo(),
		NombreApellido: "",
		DNI:            "38388388",
		Telefono:       "3493456789",
		Domicilio:      "calle falsa 123",
	}
	fmt.Print(c)

	defer fmt.Print("Ejecución terminada!")
	leerArchivo()
	c.validarDatos()
}

type cliente struct {
	Legajo         int
	NombreApellido string
	DNI            string
	Telefono       string
	Domicilio      string
}

func generarNumeroLegajo() int {
	return rand.Intn(2000)
}

func leerArchivo() string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	abirArchivo, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o esta dañado")
	}
	return string(abirArchivo)
}

func (c *cliente) validarDatos() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if c.NombreApellido == "" || c.DNI == "" || c.Telefono == "" || c.Domicilio == "" {
		panic("Se detectaron varios errores en tiempo de ejecucion")
	}
}
