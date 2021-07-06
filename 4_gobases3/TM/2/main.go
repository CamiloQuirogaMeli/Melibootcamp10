package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("productos.txt")

	if err == nil {
		fmt.Print(string(dat))
	} else {
		fmt.Println("Error leyendo archivo: ", err)
	}
}
