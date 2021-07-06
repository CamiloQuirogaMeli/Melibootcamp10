package main

import (
	"fmt"
	"os"
)

type client struct {
	fileNumber string
	completeName string
	dni string
	phoneNumber string
	address string
}

func main(){

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	readFile()
	_, err := validateClientData(
		&client{
			fileNumber:   "234lfa",
			completeName: "pepe toto",
			dni:          "",
			phoneNumber:  "424242",
			address:      "calle 1",
		})
	if err != nil {
		panic(err)
	}
}

func validateClientData(c *client) (bool, error) {
	if
	c.dni == "" || c.address == "" ||
		c.fileNumber == "" || c.dni == "" ||
		c.phoneNumber == "" {

		return false, fmt.Errorf("falta campo requerido")
	}
	return true, nil
}

func readFile(){
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.Open("customers.txt")
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o esta dañado")
	}
}

