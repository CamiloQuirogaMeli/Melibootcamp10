package main

import (
	"fmt"
	"os"
)

type Product struct {
	Id    string
	Price float64
	Cant  int
}

func main() {
	defer exeComplete()
	_, err := createFile()
	if err != nil {
		panic(err)
	}
}

func createFile() (*os.File, error) {
	filename := "inventory.csv"
	var file *os.File
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("no existe el archivo%s", filename)
	} else {
		f, e := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
		if e != nil {
			panic(e)
		} else {
			file = f
		}
		return file, nil
	}
}

func exeComplete() {
	fmt.Println("finalizando ejecución.")
}
