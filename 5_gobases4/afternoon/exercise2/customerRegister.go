package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type Customer struct {
	Legajo   int
	Name     string
	LastName string
	DNI      string
	Address  string
}

func main() {

	_, err1 := getCostumerFromFile()
	newCustomer := Customer{Legajo: 67574, Name: "Andres"}
	err2 := checkValidCustomer(newCustomer)
	defer func(err1, err2 error) {

		fmt.Println("Fin de la ejecución")
		if err1 != nil || err2 != nil {
			fmt.Println("Se detectaron varios errores en el tiempo de ejecución")
		}
		fmt.Println("No han quedado archivos abiertos")

	}(err1, err2)

}

func getCostumerFromFile() ([]byte, error) {
	defer handlePanics()
	data, err := ioutil.ReadFile("./customers.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Println(string(data))
	}
	return data, err
}
func handlePanics() {

	err := recover()
	if err != nil {
		fmt.Println(err)
	}
}

func checkValidCustomer(customer Customer) error {
	if customer.Address == "" || customer.DNI == "" || customer.LastName == "" || customer.Name == "" || customer.Legajo == 0 {
		return errors.New("existen campos vacios")
	}
	return nil
}
