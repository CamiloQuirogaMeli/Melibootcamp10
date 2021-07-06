package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type Client struct {
	id      int
	name    string
	surname string
	dni     int
	phone   int
	address string
}

func (c *Client) assignId(id int) (bool, error) {
	if id == 0 {
		return false, errors.New("El legajo no puede ser cero")
	}
	c.id = id
	return true, nil
}

func (c *Client) newClient(name, surname, address string, dni, phone int) (bool, error) {
	if name == "" || surname == "" || address == "" {
		return false, errors.New("El nombre, apellido y domicilio no pueden ser vacios")
	}
	if dni == 0 || phone == 0 {
		return false, errors.New("El dni y el telefono no pueden ser cero")
	}
	c.name = name
	c.surname = surname
	c.address = address
	c.dni = dni
	c.phone = phone
	return true, nil
}

func exist() (bool, error) {
	_, err1 := ioutil.ReadFile("./customers.txt")
	if err1 != nil {
		panic("El archivo indicado no fue encontrado o esta dañado")
	}
	return false, nil
}

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("No han quedado archivos abiertos")
	fmt.Println("Se detectaron varios errores en tiempo de ejecuión")
	fmt.Println("Fin de la Ejecución")

	exist, _ := exist()

	c := Client{}
	if !exist {
		_, err3 := c.assignId(1)
		if err3 != nil {
			panic(err3)
		}
		_, err4 := c.newClient("Karen", "Cuadrado", "9 de Julio 123", 3406459181, 35259321)
		if err4 != nil {
			panic(err4)
		}
	}
	fmt.Println(c)
}
