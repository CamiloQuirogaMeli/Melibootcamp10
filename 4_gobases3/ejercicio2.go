package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func ejercicio2() {

	archivo, err := ioutil.ReadFile("./archivo")

	if err != nil {
		log.Fatal(err)
	}

	var archivoStr []string = strings.Split(string(archivo), ",")

	if len(archivoStr)%3 == 0 {
		fmt.Printf("ID\tPrecio\tCantidad\n")
		for i := 0; i < len(archivoStr); i += 3 {
			fmt.Printf("%v\t%v\t%v\n", archivoStr[i], archivoStr[i+1], archivoStr[i+2])
		}
	} else {
		fmt.Println("Debes insertar ternas de datos")
	}

}
