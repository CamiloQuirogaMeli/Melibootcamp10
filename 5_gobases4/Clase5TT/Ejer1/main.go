package main

import (
	"fmt"
	"os"
)

func main() {
	openFile("customers.txt")

	fmt.Println("ejecucion finalizada")
}

func openFile(filename string) {
	_, err := os.Open(filename)
	if err != nil {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err)
			}
		}()
		panic("el archivo indicado no fue encontrado o esta dañado")
	}
}
