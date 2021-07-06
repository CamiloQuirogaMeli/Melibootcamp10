package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Archivo de productos \n\n")
	dat, err := ioutil.ReadFile("./archivoEjer")
	if err != nil {
		fmt.Println(err)
	} else {
		s := string(dat)
		datos := strings.Split(s, ";")
		id := "ID"
		precio := "Precio"
		cantidad := "Cantidad"
		fmt.Printf("%-6s\t\t%6s\t%6s \n", id, precio, cantidad)
		for _, v := range datos {
			v = strings.ReplaceAll(v, " ", "")
			linea := strings.Split(v, ",")
			for i := 0; i < len(linea); i++ {
				if i == 1 {
					j, _ := strconv.Atoi(linea[2])
					k, _ := strconv.Atoi(linea[i])
					fmt.Printf("%6d\t", k*j)
				} else if i == 2 {
					k, _ := strconv.Atoi(linea[i])
					fmt.Printf("%8d\t", k)
				} else {
					fmt.Print(linea[i], "\t\t")
				}
			}
			fmt.Println()
		}
	}
}
