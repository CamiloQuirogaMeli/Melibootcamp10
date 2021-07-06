package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./customers.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Println(string(data))
	}

}
