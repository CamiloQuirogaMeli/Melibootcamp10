package main

import (
	"fmt"
	"io/ioutil"
)

type product struct {
	id       int
	price    float64
	quantity int
}

func main() {

	data, err := ioutil.ReadFile("./file.txt")

	fileStr := string(data)

	fmt.Println("Id \t Price \t Quantity")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileStr)
	}

}
