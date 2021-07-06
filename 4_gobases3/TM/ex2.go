package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("ventas.dat")
	lines := strings.Split(string(dat), ";\n")
	fmt.Println("ID", "\t", "Precio\t", "\t", "Cantidad\t")
	for i, line := range lines {
		if i < len(lines)-1 {
			saleInfo := strings.Split(line, " ")
			fmt.Println(saleInfo[0], "\t\t", saleInfo[1], "\t\t", saleInfo[2])
		}
	}
}
