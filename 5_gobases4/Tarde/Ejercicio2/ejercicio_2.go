package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type cliente struct {
	legajo   int
	nombre   string
	apellido string
	dni      int
	telefono int
}

func validarDatos(c cliente) (bool, error) {
	if c.legajo != 0 && c.nombre != "" && c.apellido != "" && c.dni != 0 && c.telefono != 0 {
		return true, nil
	}
	return false, errors.New("uno o mas valores ingresados son invalidos")
}

func abrirArchivo() ([]byte, error) {
	data, err := ioutil.ReadFile("./customers.txt")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	legajo := 1234
	c := cliente{legajo, "", "", 0, 0}

	defer func() {
		errFile := recover()
		fmt.Println(errFile)

		_, errData := validarDatos(c)
		if errData != nil {
			fmt.Println(errData)
		}
		fmt.Println("Fin de la ejecucion")
		fmt.Println("Se detectaron varios errores en tiempo de ejecucion")
		fmt.Println("No han quedado archivos abiertos")

	}()

	_, err := abrirArchivo()
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o esta dañado")
	}
}
