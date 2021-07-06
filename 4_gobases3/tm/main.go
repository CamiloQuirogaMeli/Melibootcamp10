package main

import (
	"fmt"

	"github.com/lpanizza/gobases3_tm/paquetes"
)

func main() {
	path := "./file.txt"
	productos := "1,24.5,3;2,12,1;3,234.50,1"
	err := paquetes.GuardarArchivo(path, productos)

	if err != nil {
		fmt.Println(err)
	}

	data, err := paquetes.LeerArchivo(path)

	if err != nil {
		fmt.Println(err)
	}

	head, body := paquetes.ReemplazarChars(string(data))

	fmt.Print(head, body, "\n")

}
