package main

import (
	"fmt"
	"io/ioutil"
)

type Cliente struct {
	Legajo    int
	Nombre    string
	Apellido  string
	DNI       string
	Telefono  string
	Domicilio string
}

func generarLegajo(ultimoLegajo int) int {
	return ultimoLegajo + 1
}
func newCliente() *Cliente {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	cliente := Cliente{}
	ultimo, err := obtenerUltimoCliente()
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	}
	fmt.Println("Nombre: ")
	fmt.Scanln(&cliente.Nombre)
	fmt.Println("Apellido: ")
	fmt.Scanln(&cliente.Apellido)
	fmt.Println("Dni: ")
	fmt.Scanln(&cliente.DNI)
	fmt.Println("Telefono: ")
	fmt.Scanln(&cliente.Telefono)
	fmt.Println("Domicilio: ")
	fmt.Scanln(&cliente.Domicilio)
	validarCliente(cliente)
	cliente.Legajo = generarLegajo(ultimo.Legajo)
	return &cliente
}
func validarCliente(cliente Cliente) (bool, error) {
	return true, nil
}
func main() {
	fmt.Println("Nuevo cliente")
	cliente := newCliente()
	fmt.Println(cliente)
}
func obtenerUltimoCliente() (*Cliente, error) {
	cliente := Cliente{}

	_, err := ioutil.ReadFile("./customers.txt")
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}
