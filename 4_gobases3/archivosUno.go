package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	d1 := []byte("1000;20;3;\n1001;32;34;\n1002;4;56")
	err := ioutil.WriteFile("./dataProductos", d1, 0644)

	if err != nil {
		fmt.Println("No se pudo guardar el archivo correctamente")
	}

}
