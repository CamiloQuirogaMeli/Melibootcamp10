package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	customersByte, err := ioutil.ReadFile("./customers.txt")
	customers := string(customersByte)

	print(customers, err)

	fmt.Println("Ejecución finalizada")
}

func print(customers string, err error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Println(customers)
	}
}
