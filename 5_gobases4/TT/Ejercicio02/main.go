package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type Cliente struct {
	Legajo      int
	Nombre      string
	Apellido    string
	Documento   int
	NroTelefono int
	Domicilio   string
}

func (c *Cliente) legajo() {
	c.Legajo = c.Documento
}

func readFile(filename string) ([]byte, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error: el archivo indicado no fue encontrado o está dañado")
	} else {
		return data, nil
	}
}

func registrarCliente(cliente Cliente) (int, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	cliente.legajo()
	err := validarData(cliente)

	if err != nil {
		fmt.Println(err)
		return 400, errors.New("se detectaron varios errores en tiempo de ejecución")
	} else {
		return 200, nil
	}
}

func validarData(c Cliente) error {
	if c.Legajo <= 0 {
		return errors.New("legajo invalido")
	}
	if c.Nombre == "" {
		return errors.New("nombre invalido")
	}
	if c.Apellido == "" {
		return errors.New("apellido invalido")
	}
	if c.Documento == 0 {
		return errors.New("documento invalido")
	}
	if c.NroTelefono == 0 {
		return errors.New("numero de telefono invalido")
	}
	if c.Domicilio == "" {
		return errors.New("domicilio invalido")
	}
	return nil
}

func main() {
	filename := "../customers.txt"
	fileData, err := readFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(fileData))
	}

	cliente := Cliente{Apellido: "Martinez", Documento: 123456789, NroTelefono: 1134085569, Domicilio: "Triunviaro 2536"}

	status, err := registrarCliente(cliente)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(status)
	}
	fmt.Println("No han quedado archivos abiertos")
	fmt.Println("Fin de la Ejecucion")
}
