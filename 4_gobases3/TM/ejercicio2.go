package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("./go-file-products.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("ID\tPrecio\tCantidad")
	var products string = string(data)
	for i := 0; i < len(products); i++ {
		if products[i] == ';' {
			fmt.Print("\t")
		} else {
			fmt.Print(string(products[i]))
		}
	}
}
