package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Client struct {
	Legajo   string
	Name     string
	Lastname string
	Phone    string
	Addrss   string
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
		return nil, fmt.Errorf("el archivo %s no fue encontrado o está dañado”,", filename)
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
func closeFile(file *os.File) {
	fmt.Println("cerrando archivo...")
	file.Close()
}

func showFileContent(file *os.File, id string) {
	csvFile, err := os.Open("inventory.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		fmt.Println(err)
	}
	var clients []Client
	for _, line := range csvLines {
		if id == "0" {
			break
		}
		if id != line[0] {
			break
		}
		client := Client{
			Legajo:   line[0],
			Name:     line[1],
			Lastname: line[2],
			Phone:    line[3],
			Addrss:   line[4],
		}
		clients = append(clients, client)
	}

}

func exeComplete() {
	fmt.Println("fin de la ejecución.")
	fmt.Println("Se detectaron varios errores en tiempo de ejecución.")
	fmt.Println("No han quedado archivos abiertos.")
}
