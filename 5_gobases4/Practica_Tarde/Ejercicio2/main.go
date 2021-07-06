package main

import (
	"errors"
	"fmt"
	"os"
)

type Cliente struct {
	legajo    int
	nombre    string
	apellido  string
	DNI       int
	telefono  int
	domicilio string
}

func RegistrarCliente(leg int, nom string, ape string, dni int, tele int, dom string) (Cliente, error) {
	if leg == 0 {
		return Cliente{}, errors.New("El legajo no puede ser cero")
	} else if nom == "" {
		return Cliente{}, errors.New("El nombre no puede ser vacio")
	} else if ape == "" {
		return Cliente{}, errors.New("El apellido no puede ser vacio")
	} else if dni == 0 {
		return Cliente{}, errors.New("El DNI no puede ser cero")
	} else if tele == 0 {
		return Cliente{}, errors.New("El telefono no puede ser cero")
	} else if dom == "" {
		return Cliente{}, errors.New("El domicilio no puede ser vacio")
	}

	c := Cliente{leg, nom, ape, dni, tele, dom}
	return c, nil
}

func main() {

	_, err := os.Open("customers.txt")

	defer func() {
		err := recover()
		if err != nil {
			fmt.Print(err)

			c, err := RegistrarCliente(51522, "Valeria", "Macina", 92915760, 0, "Cordoba,Argentina")
			if err != nil {
				fmt.Println(err)
				fmt.Println("Se detectaron varios errores en tiempo de ejecucion")
			} else {
				fmt.Println("Se registro correctamente el cliente: ", c.apellido, ",", c.nombre)
			}
			fmt.Println("Fin de la ejecucion.")
			fmt.Println("No han quedado archivos abiertos")
		}
	}()

	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado\n")
	}
}
