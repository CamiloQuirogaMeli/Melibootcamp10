package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	totalPrecios := 0
	var precio int

	dat, err := ioutil.ReadFile("./dataProductos")

	if err == nil {

		data := strings.Split(string(dat), "\n")
		fmt.Println(data)

		fmt.Println("ID\tPrecio\tCantidad")

		for i := 0; i < len(data); i++ {
			camposItem := strings.Split(data[i], ";")
			fmt.Printf("%s\t%s\t%s\t\n", camposItem[0], camposItem[1], camposItem[2])
			precio, err = strconv.Atoi(camposItem[1])
			totalPrecios += precio
		}

		fmt.Println(totalPrecios)

	} else {
		fmt.Println("Ocurrio un error al intentar leer el archivo")
	}
}
