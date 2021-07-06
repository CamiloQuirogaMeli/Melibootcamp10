package main

import (
	"fmt"
	"os"
)

type Client struct {
	file    int
	name    string
	dni     string
	phone   int
	address string
}

func readFile(fileName string) {
	_, fileErr := os.Open("./customers.txt")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if fileErr != nil {
		panic("el archivo indicado no fue encontrado o esta roto")
	}

}

func finishingProcess() {
	fmt.Println("fin de la ejecución.")
	fmt.Println("se detectaron varios errores en tiempo de ejecución.")
	fmt.Println("no han quedado archivos abiertos.")
}

func validateFields(c Client) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover en validateFields:", r)
		}
	}()
	if c.file == 0 || c.name == "" || c.dni == "" || c.phone == 0 || c.address == "" {
		panic("hay campos del cliente que estan vacios")
	}
}

func main() {
	defer finishingProcess()
	readFile("../customers.txt")
	client := Client{}
	// client := Client{
	// 	file:    3242651324,
	// 	name:    "Carlos Infante",
	// 	dni:     "81374127",
	// 	phone:   382478921,
	// 	address: "Carrera 1 #1-1",
	// }
	validateFields(client)
}
