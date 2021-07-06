package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
}

func (e CustomError) Error() string {
	return "Se detectaron varios errores en tiempo de ejecución"
}

type Cliente struct {
	Legajo         int
	Nombre         string
	Apellido       string
	DNI            int
	NumeroTelefono int
	Domicilio      string
}

func getLegajo() (legajo int, err error) {

	legajo, err = 0, nil
	defer func() {
		recovered := recover()
		errorText := fmt.Sprint(recovered)
		err = errors.New(errorText)
	}()

	// intenta leer el archivo
	panic("error: el archivo indicado no fue encontrado o está dañado")
}

func validateCliente(dato Cliente) (clienteReturn Cliente, err error) {

	err = nil
	clienteReturn = Cliente{}

	if dato.Apellido == "" || dato.Nombre == "" || dato.Domicilio == "" || dato.DNI == 0 || dato.Legajo == 0 || dato.NumeroTelefono == 0 {
		err = CustomError{}
		return
	}
	clienteReturn = dato
	return
}

func main() {

	nombre, apellido, domicilio, dni, numero := "Nombre1", "Apellido1", "Domicilio", 34567654, 1523456789
	legajo, err := getLegajo()
	if err != nil {
		fmt.Println(err)
	} else {

		cliente := Cliente{legajo, nombre, apellido, dni, numero, domicilio}

		_, errCliente := validateCliente(cliente)

		if errCliente != nil {
			fmt.Println(errCliente)
		} else {
			fmt.Println("Se logro crear el cliente")
		}

	}
	fmt.Println("Fin de la ejecucion")
}
