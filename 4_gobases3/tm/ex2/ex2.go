package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("../ex1/product.txt")
	if err != nil {
		fmt.Println("Ocurrio un error")
	}
	fmt.Println(dat)
}
