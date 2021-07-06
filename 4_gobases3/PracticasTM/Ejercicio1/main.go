package main

import (
	"io/ioutil"
	"log"
)

func main() {
	b := []byte("ID producto;Precio;Cantidad\n1112223;300012.00;1\n444321;1000000.00;4;\n434321;50.50.00;1")
	err := ioutil.WriteFile("ejercicio1.txt", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
