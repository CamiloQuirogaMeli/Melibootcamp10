package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {

	defer func() {

		err := recover()

		if err != nil {
			fmt.Println(err)
		}

		registerClient()
		fmt.Println("End of execution")
		fmt.Println("There are no open files")

	}()

	content, err := ioutil.ReadFile("customers.txt")

	if err != nil {
		panic("error: the indicated file was not found or is corrupted")
	} else {
		fmt.Println(string(content))
	}

}

type Client struct {
	file            int
	name            string
	surname         string
	id              int
	telephoneNumber int
	address         string
}

func createClient(file int, name string, surname string, id int, telephoneNumber int, address string) (Client, error) {

	client := Client{file, name, surname, id, telephoneNumber, address}

	validatedClient, err := validateData(client)

	if err != nil {
		return validatedClient, err
	}

	return validatedClient, nil
}

func validateData(client Client) (Client, error) {

	emptyClient := Client{}

	if client.file == 0 || client.id == 0 || client.telephoneNumber == 0 {

		return emptyClient, errors.New("error: the file, ID or phone number cannot be zero, try another heat")
	}

	if client.name == "" || client.surname == "" || client.address == "" {

		return emptyClient, errors.New("error: the client cannot have empty fields")
	}

	return client, nil
}

func registerClient() {

	defer func() {

		err := recover()

		if err != nil {
			fmt.Println(err)
			fmt.Println("Multiple errors were detected at runtime")
		}

	}()

	clientToRegister, err := createClient(564312, "Joseph", "Caicedo", 0, 3118871453, "Cll 20B# 45 30")

	if err != nil {
		panic(err)
	}

	fmt.Println("Se registró el client:", clientToRegister)
}
