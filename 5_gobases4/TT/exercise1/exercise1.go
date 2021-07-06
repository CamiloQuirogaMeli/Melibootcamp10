package main

import (
	"fmt"
	"os"
)

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

func main() {
	readFile("../customers.txt")
	fmt.Println("Ejecucion finalizada")
}
