package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("file.txt")

	defer func() {
		fmt.Println("ejecución finalizada")
	}()

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}
