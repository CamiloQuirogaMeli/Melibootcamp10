package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Iniciando programa")
	_, err := os.Open("customers.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Ejecucion finalizada")
}
