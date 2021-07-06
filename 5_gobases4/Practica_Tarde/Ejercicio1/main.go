package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("customers.txt")

	defer func() {
		fmt.Println("ejecucion finalizada")
	}()

	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}

}
