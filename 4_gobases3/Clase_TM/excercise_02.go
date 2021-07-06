package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./dat1")

	if err != nil {
		fmt.Println(err)
	}

	dataString := string(data)
	fmt.Println(dataString)

	// falta sumar los monto totales
}
