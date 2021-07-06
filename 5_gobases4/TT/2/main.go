package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	var leg int
	fmt.Println("Ingrese el legajo:")
	fmt.Scanln(&leg)
	getCustomers()

	var nombre string
	var dni string
	var tel string
	var domicilio string
	fmt.Println("Ingrese el nombre:")
	fmt.Scanln(&nombre)
	fmt.Println("Ingrese el dni:")
	fmt.Scanln(&dni)
	fmt.Println("Ingrese el telefono:")
	fmt.Scanln(&tel)
	fmt.Println("Ingrese el domicilio:")
	fmt.Scanln(&domicilio)

	_, err := validateCustomer(leg, nombre, dni, tel, domicilio)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Customer agregado")
	}

	fmt.Println("Fin de la ejecucion")
	fmt.Println("Se detectaron varios errores en tiempo de ejecucion")
	fmt.Println("No han quedado archivos abiertos")

}

func getCustomers() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := ioutil.ReadFile("customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}
}

func validateCustomer(leg int, nom string, dni string, tel string, dom string) (Customer, error) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	switch {
	case leg == 0:
		return Customer{}, errors.New("error: el legajo ingresado es cero")
	case nom == "":
		return Customer{}, errors.New("error: el nombre ingresado es invalido")
	case dni == "":
		return Customer{}, errors.New("error: no ingreso numero de documento")
	case tel == "":
		return Customer{}, errors.New("error: no ingreso telefono")
	case dom == "":
		return Customer{}, errors.New("error: no ingreso domicilio")
	default:
		return Customer{legajo: leg, nombre: nom, dni: dni, telefono: tel, domicilio: dom}, nil
	}
}

type Customer struct {
	legajo    int
	nombre    string
	dni       string
	telefono  string
	domicilio string
}
