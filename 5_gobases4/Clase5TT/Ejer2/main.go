package main

import (
	"fmt"
	"os"
)

type user struct {
	Legajo    int64
	Nombre    string
	Apellido  string
	DNI       int64
	Numero    int64
	Domicilio string
}

func main() {
	var legajo int64 = 3
	openFile("customers.txt")
	createUser(legajo, "Tom", "Sawyer", 0, 3654343, "Avenida siempre viva")
	fmt.Println("ejecucion finalizada")
	fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	fmt.Println("No han quedado archivos abiertos")
}
func createUser(legajo int64, name, lastname string, DNI, number int64, address string) (user, error) {
	var usr user
	if legajo <= 0 || name == "" || lastname == "" || DNI <= 0 || number <= 0 || address == "" {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err)
			}
		}()
		panic("error: no todos los datos son validos")
	} else {
		usr = user{legajo, name, lastname, DNI, number, address}
		return usr, nil
	}
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
