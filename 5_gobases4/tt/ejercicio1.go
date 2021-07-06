package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	_, err := ioutil.ReadFile("customers.txt")

	defer func() {fmt.Println("ejecucion finalizada")}()
	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}

}
