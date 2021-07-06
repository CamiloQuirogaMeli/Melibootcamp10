package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	byteProductos, err := ioutil.ReadFile("./productos.txt")
	if err == nil {
		productos := string(byteProductos)
		fmt.Printf("ID\t\t\tPrecio\tCantidad\n")
		for _, caracter := range productos {
			switch caracter {
			case ';':
				fmt.Printf("\n")
			case ',':
				fmt.Printf("\t")
			default:
				fmt.Printf("%c", caracter)
			}
		}
	} else {
		fmt.Print("Error: ", err)
	}
}
