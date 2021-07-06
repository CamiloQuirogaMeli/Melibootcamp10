package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	defer func() {
		fmt.Println("Ejecucion finalizada")
	}()

	_, err := ioutil.ReadFile("./customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	fmt.Println("Ejecucion finalizada")

}
